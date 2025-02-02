package findingapistructs

type FindItemsAdvancedRequest struct {
	XMLName           string          `xml:"findItemsAdvancedRequest"`
	XMLNS             string          `xml:"xmlns,attr"`
	AspectFilters     []AspectFilter  `xml:"aspectFilter"`
	CategoryIds       []string        `xml:"categoryId"`
	DescriptionSearch bool            `xml:"descriptionSearch"`
	DomainFilters     []DomainFilter  `xml:"domainFilter"`
	ItemFilters       []ItemFilter    `xml:"itemFilter"`
	Keywords          string          `xml:"keywords"`
	OutputSelectors   []string        `xml:"outputSelector"`
	Affiliate         Affiliate       `xml:"affiliate"`
	BuyerPostalCode   string          `xml:"buyerPostalCode"`
	PaginationInput   PaginationInput `xml:"paginationInput"`
	SortOrder         string          `xml:"sortOrder"`
}

type AspectFilter struct {
	AspectName       string   `xml:"aspectName"`
	AspectValueNames []string `xml:"aspectValueName"`
}

type DomainFilter struct {
	DomainNames []string `xml:"domainName"`
}

type ItemFilter struct {
	Name       string   `xml:"name"`
	ParamName  string   `xml:"paramName"`
	ParamValue string   `xml:"paramValue"`
	Values     []string `xml:"value"`
}

type Affiliate struct {
	CustomID     string `xml:"customId"`
	GeoTargeting bool   `xml:"geoTargeting"`
	NetworkID    string `xml:"networkId"`
	TrackingID   string `xml:"trackingId"`
}

type PaginationInput struct {
	EntriesPerPage int `xml:"entriesPerPage"`
	PageNumber     int `xml:"pageNumber"`
}
