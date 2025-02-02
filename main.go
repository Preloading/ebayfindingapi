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

	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read response body")
	}
	return c.SendString(string(body))
}
