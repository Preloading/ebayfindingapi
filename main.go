package main

import (
	"bytes"
	"ebaysearchapi/findingapistructs"
	"ebaysearchapi/modernapistructs"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	// AppID is the eBay application ID
	AppID                 = "JohnFort-echobayf-PRD-266f79d08-ec4532d2"
	CertId                = "PRD-66f79d08b106-f1c5-448f-8a1f-b688"
	AccessToken           string
	AccessTokenExpiration time.Time
)

func main() {
	GenerateNewAccessToken()

	app := fiber.New()

	app.Post("/services/search/FindingService/v1", helloWorld)

	app.Listen(":3000")
}

func GenerateNewAccessToken() {
	client := &http.Client{}
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "https://api.ebay.com/oauth/api_scope")

	req, err := http.NewRequest("POST", "https://api.ebay.com/identity/v1/oauth2/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to create AUTH request: %v", err))
		return
	}

	// Fix credentials encoding
	credentials := base64.StdEncoding.EncodeToString([]byte(AppID + ":" + CertId))
	req.Header.Set("Authorization", "Basic "+credentials)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send AUTH request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// // Read and print the response body
	// responseBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Failed to read AUTH response body: %v\n", err)
	// 	return
	// }
	// fmt.Printf("AUTH response body: %s\n", string(responseBody))

	// Decode the response using the copied body
	var authResponse modernapistructs.OAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		fmt.Printf("Failed to decode AUTH response: %v\n", err)
		return
	}

	// Save the access token and expiration time
	AccessToken = authResponse.AccessToken
	AccessTokenExpiration = time.Now().Add(time.Duration(authResponse.ExpiresIn) * time.Second)
}

func helloWorld(c *fiber.Ctx) error {
	//app_name := c.Get("X-EBAY-SOA-SECURITY-APPNAME")
	return FindItemsAdvancedRequest(c)
}

func FindItemsAdvancedRequest(c *fiber.Ctx) error {
	var request findingapistructs.FindItemsAdvancedRequest
	err := xml.Unmarshal(c.Body(), &request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request")
	}

	// https://developer.ebay.com/api-docs/buy/browse/resources/item_summary/methods/search
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.ebay.com/buy/browse/v1/item_summary/search"+

		"?q="+url.QueryEscape(request.Keywords)+
		"&limit="+url.QueryEscape(fmt.Sprintf("%d", request.PaginationInput.EntriesPerPage))+
		"&offset="+url.QueryEscape(fmt.Sprintf("%d", request.PaginationInput.EntriesPerPage*(request.PaginationInput.PageNumber-1))),

		nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create request")
	}
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	fmt.Println(c.Get("X-EBAY-SOA-GLOBAL-ID"))
	req.Header.Add("X-EBAY-C-MARKETPLACE-ID", c.Get("X-EBAY-SOA-GLOBAL-ID"))

	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read response body")
	}

	// Decode the response using the copied body
	var searchResp modernapistructs.SearchPagedCollection
	err = json.Unmarshal(body, &searchResp)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to decode response")
	}

	items := []findingapistructs.Item{}

	for _, itemSummary := range searchResp.ItemSummaries {
		item := findingapistructs.Item{
			ItemId:   itemSummary.LegacyItemId,
			Title:    itemSummary.Title,
			GlobalId: itemSummary.ListingMarketplaceId,
			PrimaryCategory: findingapistructs.Category{
				CategoryId:   itemSummary.Categories[0].CategoryId,
				CategoryName: itemSummary.Categories[0].CategoryName,
			},
			SecondaryCategory: func() findingapistructs.Category {
				if len(itemSummary.Categories) < 2 {
					return findingapistructs.Category{
						CategoryId:   itemSummary.Categories[1].CategoryId,
						CategoryName: itemSummary.Categories[1].CategoryName,
					}
				}
				return findingapistructs.Category{}
			}(),
			GalleryURL:  itemSummary.Image.ImageUrl,
			ViewItemURL: itemSummary.ItemWebUrl,
			AutoPay:     false, // figue dis out
			PostalCode:  itemSummary.ItemLocation.PostalCode,
			//Location:    itemSummary.ItemLocation.City, // City might be unavailible
			Country: itemSummary.ItemLocation.Country,
			ShippingInfo: findingapistructs.ShippingInfo{ // seems kinda finicky
				ShippingType: func() string { // This could be done better
					if len(itemSummary.ShippingOptions) == 0 {
						return "NotSpecified"
					}
					if itemSummary.ShippingOptions[0].ShippingCost.Currency == "" {
						return "Free"
					}
					if itemSummary.ShippingOptions[0].ShippingCostType == "FLAT" {
						return "Flat"
					}
					if itemSummary.ShippingOptions[0].ShippingCostType == "CALCULATED" {
						return "Calculated"
					}
					return "NotSpecified"
				}(),
				ShippingServiceCost: findingapistructs.Amount{
					CurrencyId: itemSummary.ShippingOptions[0].ShippingCost.Currency,
					Value:      itemSummary.ShippingOptions[0].ShippingCost.Value,
				},
			},
			SellingStatus: findingapistructs.SellingStatus{
				BidCount: itemSummary.BidCount,
				CurrentPrice: findingapistructs.Amount{
					CurrencyId: itemSummary.Price.Currency,
					Value:      itemSummary.Price.Value,
				},
				ConvertedCurrentPrice: func() findingapistructs.Amount {
					if itemSummary.Price.ConvertedFromCurrency == "" {
						return findingapistructs.Amount{
							CurrencyId: itemSummary.Price.Currency,
							Value:      itemSummary.Price.Value,
						}
					}
					return findingapistructs.Amount{
						CurrencyId: itemSummary.Price.ConvertedFromCurrency,
						Value:      itemSummary.Price.ConvertedFromValue,
					}
				}(),
				SellingState: "Active", // cheat
				TimeLeft: func() string {
					var endTime time.Time
					if itemSummary.ItemEndDate != "" {
						endTime, _ = time.Parse(time.RFC3339, itemSummary.ItemEndDate)
					} else {
						// Use the same logic as in ListingInfo
						startTime, _ := time.Parse(time.RFC3339, itemSummary.ItemCreationDate)
						currentTime := time.Now()
						listingStart := startTime

						for {
							nextMonth := listingStart.AddDate(0, 1, 0)
							if nextMonth.After(currentTime) {
								break
							}
							listingStart = nextMonth
						}
						endTime = listingStart.AddDate(0, 1, 0).Add(-10 * time.Second)
					}

					// If the item has ended, return P0DT0H0M0S
					if time.Now().After(endTime) {
						return "P0DT0H0M0S"
					}

					// Calculate duration between now and end time
					duration := endTime.Sub(time.Now())
					days := int(duration.Hours() / 24)
					hours := int(duration.Hours()) % 24
					minutes := int(duration.Minutes()) % 60
					seconds := int(duration.Seconds()) % 60

					return fmt.Sprintf("P%dDT%dH%dM%dS", days, hours, minutes, seconds)
				}(),
			},
			ListingInfo: findingapistructs.ListingInfo{
				BuyItNowAvailable: (Contains(itemSummary.BuyingOptions, "AUCTION") || Contains(itemSummary.BuyingOptions, "BEST_OFFER")) && Contains(itemSummary.BuyingOptions, "FIXED_PRICE"),
				BuyItNowPrice: func() findingapistructs.Amount {
					if (Contains(itemSummary.BuyingOptions, "AUCTION") || Contains(itemSummary.BuyingOptions, "BEST_OFFER")) && Contains(itemSummary.BuyingOptions, "FIXED_PRICE") {
						return findingapistructs.Amount{
							CurrencyId: itemSummary.Price.Currency,
							Value:      itemSummary.Price.Value,
						}
					}
					return findingapistructs.Amount{}
				}(),
				ConvertedBuyItNowPrice: func() findingapistructs.Amount {
					if (Contains(itemSummary.BuyingOptions, "AUCTION") || Contains(itemSummary.BuyingOptions, "BEST_OFFER")) && Contains(itemSummary.BuyingOptions, "FIXED_PRICE") {
						if itemSummary.Price.ConvertedFromCurrency == "" {
							return findingapistructs.Amount{
								CurrencyId: itemSummary.Price.Currency,
								Value:      itemSummary.Price.Value,
							}
						}
						return findingapistructs.Amount{
							CurrencyId: itemSummary.Price.ConvertedFromCurrency,
							Value:      itemSummary.Price.ConvertedFromValue,
						}
					}
					return findingapistructs.Amount{}
				}(),
				BestOfferEnabled: Contains(itemSummary.BuyingOptions, "BEST_OFFER"),
				ListingType: func() string {
					buyItNow := false
					if Contains(itemSummary.BuyingOptions, "FIXED_PRICE") {
						buyItNow = true
					}
					if (Contains(itemSummary.BuyingOptions, "AUCTION") || Contains(itemSummary.BuyingOptions, "BEST_OFFER")) && buyItNow {
						return "AuctionWithBIN"
					}
					if buyItNow {
						return "FixedPrice"
					}
					if Contains(itemSummary.BuyingOptions, "AUCTION") {
						return "Auction"
					}
					if Contains(itemSummary.BuyingOptions, "BEST_OFFER") {
						return "Auction"
					}
					if Contains(itemSummary.BuyingOptions, "CLASSIFIED") {
						return "Classified"
					}
					return "Unknown"
				}(),
				StartTime: func() time.Time {
					// If we have an end time, use the actual item creation date
					if itemSummary.ItemEndDate != "" {
						t, _ := time.Parse(time.RFC3339, itemSummary.ItemCreationDate)
						return t
					}

					// Otherwise, start from item creation and increment months until we'd exceed current time
					creationTime, _ := time.Parse(time.RFC3339, itemSummary.ItemCreationDate)
					currentTime := time.Now()
					startTime := creationTime

					for {
						nextMonth := startTime.AddDate(0, 1, 0)
						if nextMonth.After(currentTime) {
							break
						}
						startTime = nextMonth
					}
					return startTime
				}(),
				EndTime: func() time.Time {
					// If we have an end time, use it
					if itemSummary.ItemEndDate != "" {
						t, _ := time.Parse(time.RFC3339, itemSummary.ItemEndDate)
						return t
					}

					// Otherwise calculate based on the start time we determined
					startTime, _ := time.Parse(time.RFC3339, itemSummary.ItemCreationDate)
					currentTime := time.Now()
					listingStart := startTime

					// Find the correct start time first
					for {
						nextMonth := listingStart.AddDate(0, 1, 0)
						if nextMonth.After(currentTime) {
							break
						}
						listingStart = nextMonth
					}

					// Return start time + 1 month - 10 seconds
					return listingStart.AddDate(0, 1, 0).Add(-10 * time.Second)
				}(),
				Gift: false,
			},
			SellerInfo: findingapistructs.SellerInfo{
				FeedbackScore:           int64(itemSummary.Seller.FeedbackScore),
				PositiveFeedbackPercent: itemSummary.Seller.FeedbackPercentage,
				SellerUserName:          itemSummary.Seller.Username,
				TopRatedSeller:          false, // no one is at the top.
				FeedbackRatingStar: func() string {
					score := itemSummary.Seller.FeedbackScore
					switch {
					case score >= 1000000:
						return "SilverShooting"
					case score >= 500000:
						return "GreenShooting"
					case score >= 100000:
						return "RedShooting"
					case score >= 50000:
						return "PurpleShooting"
					case score >= 25000:
						return "TurquoiseShooting"
					case score >= 10000:
						return "YellowShooting"
					case score >= 5000:
						return "Green"
					case score >= 1000:
						return "Red"
					case score >= 500:
						return "Purple"
					case score >= 100:
						return "Turquoise"
					case score >= 50:
						return "Blue"
					case score >= 10:
						return "Yellow"
					default:
						return "None"
					}
				}(),
			},
		}
		items = append(items, item)
	}

	findingResp := findingapistructs.FindItemsAdvancedResponse{
		Xmlns:     "http://www.ebay.com/marketplace/search/v1/services",
		Ack:       "Success",
		Version:   "1.13.0",
		Timestamp: time.Now(),
		SearchResult: findingapistructs.SearchResult{
			Count: len(searchResp.ItemSummaries),
			Items: items,
		},
		PaginationOutput: findingapistructs.PaginationOutput{
			EntriesPerPage: request.PaginationInput.EntriesPerPage,
			PageNumber:     request.PaginationInput.PageNumber,
			TotalEntries:   searchResp.Total,
			TotalPages:     searchResp.Total / request.PaginationInput.EntriesPerPage,
		},
	}

	// Encode the response into xml
	result, err := xml.Marshal(findingResp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to encode response")
	}

	c.Set("Content-Type", "application/xml")
	return c.SendString("<?xml version='1.0' encoding='UTF-8'?>\n" + string(result))

	//return c.SendString(string(body))
}

// helper function courtesy of Github Copilot
func Contains[T comparable](arr []T, element T) bool {
	for _, v := range arr {
		if v == element {
			return true
		}
	}
	return false
}
