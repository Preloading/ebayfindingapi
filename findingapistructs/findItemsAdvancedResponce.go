package findingapistructs

import "time"

type FindItemsAdvancedResponse struct {
	XMLName                     string                      `xml:"findItemsAdvancedResponse"`
	Xmlns                       string                      `xml:"xmlns,attr"`
	Ack                         string                      `xml:"ack"`
	Version                     string                      `xml:"version"`
	Timestamp                   time.Time                   `xml:"timestamp"`
	SearchResult                SearchResult                `xml:"searchResult"`
	PaginationOutput            PaginationOutput            `xml:"paginationOutput"`
	ItemSearchURL               string                      `xml:"itemSearchURL"`
	CategoryHistogramContainer  CategoryHistogramContainer  `xml:"categoryHistogramContainer"`
	AspectHistogramContainer    AspectHistogramContainer    `xml:"aspectHistogramContainer"`
	ConditionHistogramContainer ConditionHistogramContainer `xml:"conditionHistogramContainer,omitempty"`
	ErrorMessage                ErrorMessage                `xml:"errorMessage,omitempty"`
}

type AspectHistogramContainer struct {
	Aspects           []Aspect `xml:"aspect"`
	DomainDisplayName string   `xml:"domainDisplayName"`
	DomainName        string   `xml:"domainName"`
}

type Aspect struct {
	Name            string           `xml:"name,attr"`
	ValueHistograms []ValueHistogram `xml:"valueHistogram"`
}

type ValueHistogram struct {
	ValueName string `xml:"valueName,attr"`
	Count     int64  `xml:"count"`
}

type CategoryHistogramContainer struct {
	CategoryHistogram []CategoryHistogram `xml:"categoryHistogram"`
}

type CategoryHistogram struct {
	CategoryId             string              `xml:"categoryId"`
	CategoryName           string              `xml:"categoryName"`
	ChildCategoryHistogram []CategoryHistogram `xml:"childCategoryHistogram"`
	Count                  int64               `xml:"count"`
}

type ConditionHistogramContainer struct {
	ConditionHistogram []ConditionHistogram `xml:"conditionHistogram"`
}

type ConditionHistogram struct {
	Condition Condition `xml:"condition"`
	Count     int       `xml:"count"`
}

type Condition struct {
	ConditionDisplayName string `xml:"conditionDisplayName"`
	ConditionId          int    `xml:"conditionId"`
}

type ErrorMessage struct {
	Errors []ErrorData `xml:"error"`
}

type ErrorData struct {
	Category    string           `xml:"category"`
	Domain      string           `xml:"domain"`
	ErrorId     int64            `xml:"errorId"`
	ExceptionId string           `xml:"exceptionId"`
	Message     string           `xml:"message"`
	Parameters  []ErrorParameter `xml:"parameter"`
	Severity    string           `xml:"severity"`
	Subdomain   string           `xml:"subdomain"`
}

type ErrorParameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type PaginationOutput struct {
	EntriesPerPage int `xml:"entriesPerPage"`
	PageNumber     int `xml:"pageNumber"`
	TotalEntries   int `xml:"totalEntries"`
	TotalPages     int `xml:"totalPages"`
}

type SearchResult struct {
	Count int    `xml:"count,attr"`
	Items []Item `xml:"item"`
}

type Item struct {
	ItemId          string        `xml:"itemId"`
	Title           string        `xml:"title"`
	GlobalId        string        `xml:"globalId"`
	PrimaryCategory Category      `xml:"primaryCategory"`
	GalleryURL      string        `xml:"galleryURL"`
	ViewItemURL     string        `xml:"viewItemURL,omitempty"`
	AutoPay         bool          `xml:"autoPay"`
	PostalCode      string        `xml:"postalCode"`
	Location        string        `xml:"location"`
	Country         string        `xml:"country"`
	ShippingInfo    ShippingInfo  `xml:"shippingInfo"`
	SellingStatus   SellingStatus `xml:"sellingStatus"`
	ListingInfo     ListingInfo   `xml:"listingInfo"`
	Condition       Condition     `xml:"condition"`
	TopRatedListing bool          `xml:"topRatedListing"`

	CharityId            string               `xml:"charityId,omitempty"`
	Compatibility        string               `xml:"compatibility,omitempty"` // only for vehicle parts and accessories
	Distance             Distance             `xml:"distance,omitempty"`
	EBayPlusEnabled      bool                 `xml:"eBayPlusEnabled,omitempty"`
	GalleryInfoContainer GalleryInfoContainer `xml:"galleryInfoContainer,omitempty"`
	PaymentMethods       []string             `xml:"paymentMethod,omitempty"`
	PictureURLLarge      string               `xml:"pictureURLLarge,omitempty"`
	PictureURLSuperSize  string               `xml:"pictureURLSuperSize,omitempty"`
	ReturnsAccepted      bool                 `xml:"returnsAccepted,omitempty"`
	SecondaryCategory    Category             `xml:"secondaryCategory,omitempty"`
	SellerInfo           SellerInfo           `xml:"sellerInfo,omitempty"`
	StoreInfo            Storefront           `xml:"storeInfo,omitempty"`
	Subtitle             string               `xml:"subtitle,omitempty"`
}

type Distance struct {
	Unit  string  `xml:"unit,attr"`
	Value float64 `xml:",chardata"`
}

type GalleryInfoContainer struct {
	GalleryURLs []GalleryURL `xml:"galleryURL"`
}

type GalleryURL struct {
	Size  string `xml:"gallerySize,attr"`
	Value string `xml:",chardata"`
}

type ListingInfo struct {
	BestOfferEnabled       bool      `xml:"bestOfferEnabled"`
	BuyItNowAvailable      bool      `xml:"buyItNowAvailable"`
	BuyItNowPrice          Amount    `xml:"buyItNowPrice,omitempty"`
	ConvertedBuyItNowPrice Amount    `xml:"convertedBuyItNowPrice,omitempty"`
	EndTime                time.Time `xml:"endTime"`
	Gift                   bool      `xml:"gift"`
	ListingType            string    `xml:"listingType"`
	StartTime              time.Time `xml:"startTime"`
	WatchCount             int       `xml:"watchCount,omitempty"`
}

type Amount struct {
	CurrencyId string `xml:"currencyId,attr"`
	Value      string `xml:",chardata"`
}

type Category struct {
	CategoryId   string `xml:"categoryId"`
	CategoryName string `xml:"categoryName"`
}

type SellerInfo struct {
	FeedbackRatingStar      string `xml:"feedbackRatingStar"`
	FeedbackScore           int64  `xml:"feedbackScore"`
	PositiveFeedbackPercent string `xml:"positiveFeedbackPercent"`
	SellerUserName          string `xml:"sellerUserName"`
	TopRatedSeller          bool   `xml:"topRatedSeller"`
}

type SellingStatus struct {
	BidCount              int    `xml:"bidCount,omitempty"`
	ConvertedCurrentPrice Amount `xml:"convertedCurrentPrice"`
	CurrentPrice          Amount `xml:"currentPrice"`
	SellingState          string `xml:"sellingState"`
	TimeLeft              string `xml:"timeLeft"`
}

type ShippingInfo struct {
	ExpeditedShipping       bool     `xml:"expeditedShipping,omitempty"`
	HandlingTime            int      `xml:"handlingTime,omitempty"`
	OneDayShippingAvailable bool     `xml:"oneDayShippingAvailable,omitempty"`
	ShippingServiceCost     Amount   `xml:"shippingServiceCost,omitempty"`
	ShippingType            string   `xml:"shippingType"`
	ShipToLocations         []string `xml:"shipToLocations"`
}

type Storefront struct {
	StoreName string `xml:"storeName"`
	StoreURL  string `xml:"storeURL,omitempty"`
}
