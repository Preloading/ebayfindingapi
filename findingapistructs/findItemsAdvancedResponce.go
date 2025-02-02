package findingapistructs

import "time"

type FindItemsAdvancedResponse struct {
	XMLName                     string                      `xml:"findItemsAdvancedResponse"`
	Xmlns                       string                      `xml:"xmlns,attr"`
	AspectHistogramContainer    AspectHistogramContainer    `xml:"aspectHistogramContainer"`
	CategoryHistogramContainer  CategoryHistogramContainer  `xml:"categoryHistogramContainer"`
	ConditionHistogramContainer ConditionHistogramContainer `xml:"conditionHistogramContainer"`
	Ack                         string                      `xml:"ack"`
	ErrorMessage                ErrorMessage                `xml:"errorMessage"`
	ItemSearchURL               string                      `xml:"itemSearchURL"`
	PaginationOutput            PaginationOutput            `xml:"paginationOutput"`
	SearchResult                SearchResult                `xml:"searchResult"`
	Timestamp                   time.Time                   `xml:"timestamp"`
	Version                     string                      `xml:"version"`
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
	AutoPay              bool                 `xml:"autoPay"`
	CharityId            string               `xml:"charityId"`
	Compatibility        string               `xml:"compatibility"`
	Condition            Condition            `xml:"condition"`
	Country              string               `xml:"country"`
	Distance             Distance             `xml:"distance"`
	EBayPlusEnabled      bool                 `xml:"eBayPlusEnabled"`
	GalleryInfoContainer GalleryInfoContainer `xml:"galleryInfoContainer"`
	GalleryURL           string               `xml:"galleryURL"`
	GlobalId             string               `xml:"globalId"`
	ItemId               string               `xml:"itemId"`
	ListingInfo          ListingInfo          `xml:"listingInfo"`
	Location             string               `xml:"location"`
	PaymentMethods       []string             `xml:"paymentMethod"`
	PictureURLLarge      string               `xml:"pictureURLLarge"`
	PictureURLSuperSize  string               `xml:"pictureURLSuperSize"`
	PostalCode           string               `xml:"postalCode"`
	PrimaryCategory      Category             `xml:"primaryCategory"`
	ReturnsAccepted      bool                 `xml:"returnsAccepted"`
	SecondaryCategory    Category             `xml:"secondaryCategory"`
	SellerInfo           SellerInfo           `xml:"sellerInfo"`
	SellingStatus        SellingStatus        `xml:"sellingStatus"`
	ShippingInfo         ShippingInfo         `xml:"shippingInfo"`
	StoreInfo            Storefront           `xml:"storeInfo"`
	Subtitle             string               `xml:"subtitle"`
	Title                string               `xml:"title"`
	TopRatedListing      bool                 `xml:"topRatedListing"`
	ViewItemURL          string               `xml:"viewItemURL"`
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
	BuyItNowPrice          Amount    `xml:"buyItNowPrice"`
	ConvertedBuyItNowPrice Amount    `xml:"convertedBuyItNowPrice"`
	EndTime                time.Time `xml:"endTime"`
	Gift                   bool      `xml:"gift"`
	ListingType            string    `xml:"listingType"`
	StartTime              time.Time `xml:"startTime"`
	WatchCount             int       `xml:"watchCount"`
}

type Amount struct {
	CurrencyId string  `xml:"currencyId,attr"`
	Value      float64 `xml:",chardata"`
}

type Category struct {
	CategoryId   string `xml:"categoryId"`
	CategoryName string `xml:"categoryName"`
}

type SellerInfo struct {
	FeedbackRatingStar      string  `xml:"feedbackRatingStar"`
	FeedbackScore           int64   `xml:"feedbackScore"`
	PositiveFeedbackPercent float64 `xml:"positiveFeedbackPercent"`
	SellerUserName          string  `xml:"sellerUserName"`
	TopRatedSeller          bool    `xml:"topRatedSeller"`
}

type SellingStatus struct {
	BidCount              int    `xml:"bidCount"`
	ConvertedCurrentPrice Amount `xml:"convertedCurrentPrice"`
	CurrentPrice          Amount `xml:"currentPrice"`
	SellingState          string `xml:"sellingState"`
	TimeLeft              string `xml:"timeLeft"`
}

type ShippingInfo struct {
	ExpeditedShipping       bool     `xml:"expeditedShipping"`
	HandlingTime            int      `xml:"handlingTime"`
	OneDayShippingAvailable bool     `xml:"oneDayShippingAvailable"`
	ShippingServiceCost     Amount   `xml:"shippingServiceCost"`
	ShippingType            string   `xml:"shippingType"`
	ShipToLocations         []string `xml:"shipToLocations"`
}

type Storefront struct {
	StoreName string `xml:"storeName"`
	StoreURL  string `xml:"storeURL"`
}
