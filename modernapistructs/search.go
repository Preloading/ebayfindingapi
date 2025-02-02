package modernapistructs

type SearchPagedCollection struct {
	AutoCorrections AutoCorrections `json:"autoCorrections"`
	Href            string          `json:"href"`
	ItemSummaries   []ItemSummary   `json:"itemSummaries"`
	Limit           int             `json:"limit"`
	Next            string          `json:"next"`
	Offset          int             `json:"offset"`
	Prev            string          `json:"prev"`
	Refinement      Refinement      `json:"refinement"`
	Total           int             `json:"total"`
	Warnings        []ErrorDetailV3 `json:"warnings"`
}

type AutoCorrections struct {
	Q string `json:"q"`
}

type ItemSummary struct {
	AdditionalImages           []Image                 `json:"additionalImages"`
	AdultOnly                  bool                    `json:"adultOnly"`
	AvailableCoupons           bool                    `json:"availableCoupons"`
	BidCount                   int                     `json:"bidCount"`
	BuyingOptions              []string                `json:"buyingOptions"`
	Categories                 []Category              `json:"categories"`
	CompatibilityMatch         string                  `json:"compatibilityMatch"`
	CompatibilityProperties    []CompatibilityProperty `json:"compatibilityProperties"`
	Condition                  string                  `json:"condition"`
	ConditionId                string                  `json:"conditionId"`
	CurrentBidPrice            ConvertedAmount         `json:"currentBidPrice"`
	DistanceFromPickupLocation TargetLocation          `json:"distanceFromPickupLocation"`
	EnergyEfficiencyClass      string                  `json:"energyEfficiencyClass"`
	Epid                       string                  `json:"epid"`
	Image                      Image                   `json:"image"`
	ItemAffiliateWebUrl        string                  `json:"itemAffiliateWebUrl"`
	ItemCreationDate           string                  `json:"itemCreationDate"`
	ItemEndDate                string                  `json:"itemEndDate"`
	ItemGroupHref              string                  `json:"itemGroupHref"`
	ItemGroupType              string                  `json:"itemGroupType"`
	ItemHref                   string                  `json:"itemHref"`
	ItemId                     string                  `json:"itemId"`
	ItemLocation               ItemLocationImpl        `json:"itemLocation"`
	ItemWebUrl                 string                  `json:"itemWebUrl"`
	LeafCategoryIds            []string                `json:"leafCategoryIds"`
	LegacyItemId               string                  `json:"legacyItemId"`
	ListingMarketplaceId       string                  `json:"listingMarketplaceId"`
	MarketingPrice             MarketingPrice          `json:"marketingPrice"`
	PickupOptions              []PickupOptionSummary   `json:"pickupOptions"`
	Price                      ConvertedAmount         `json:"price"`
	PriceDisplayCondition      string                  `json:"priceDisplayCondition"`
	PriorityListing            bool                    `json:"priorityListing"`
	QualifiedPrograms          []string                `json:"qualifiedPrograms"`
	Seller                     Seller                  `json:"seller"`
	ShippingOptions            []ShippingOptionSummary `json:"shippingOptions"`
	ShortDescription           string                  `json:"shortDescription"`
	ThumbnailImages            []Image                 `json:"thumbnailImages"`
	Title                      string                  `json:"title"`
	TopRatedBuyingExperience   bool                    `json:"topRatedBuyingExperience"`
	TyreLabelImageUrl          string                  `json:"tyreLabelImageUrl"`
	UnitPrice                  ConvertedAmount         `json:"unitPrice"`
	UnitPricingMeasure         string                  `json:"unitPricingMeasure"`
	WatchCount                 int                     `json:"watchCount"`
}

type Image struct {
	Height   int    `json:"height"`
	ImageUrl string `json:"imageUrl"`
	Width    int    `json:"width"`
}

type Category struct {
	CategoryId   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

type CompatibilityProperty struct {
	LocalizedName string `json:"localizedName"`
	Name          string `json:"name"`
	Value         string `json:"value"`
}

type ConvertedAmount struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}

type TargetLocation struct {
	UnitOfMeasure string `json:"unitOfMeasure"`
	Value         string `json:"value"`
}

type ItemLocationImpl struct {
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	City            string `json:"city"`
	Country         string `json:"country"`
	County          string `json:"county"`
	PostalCode      string `json:"postalCode"`
	StateOrProvince string `json:"stateOrProvince"`
}

type MarketingPrice struct {
	DiscountAmount     ConvertedAmount `json:"discountAmount"`
	DiscountPercentage string          `json:"discountPercentage"`
	OriginalPrice      ConvertedAmount `json:"originalPrice"`
	PriceTreatment     string          `json:"priceTreatment"`
}

type PickupOptionSummary struct {
	PickupLocationType string `json:"pickupLocationType"`
}

type Seller struct {
	FeedbackPercentage string `json:"feedbackPercentage"`
	FeedbackScore      int    `json:"feedbackScore"`
	SellerAccountType  string `json:"sellerAccountType"`
	Username           string `json:"username"`
}

type ShippingOptionSummary struct {
	GuaranteedDelivery       bool            `json:"guaranteedDelivery"`
	MaxEstimatedDeliveryDate string          `json:"maxEstimatedDeliveryDate"`
	MinEstimatedDeliveryDate string          `json:"minEstimatedDeliveryDate"`
	ShippingCost             ConvertedAmount `json:"shippingCost"`
	ShippingCostType         string          `json:"shippingCostType"`
}

type Refinement struct {
	AspectDistributions       []AspectDistribution       `json:"aspectDistributions"`
	BuyingOptionDistributions []BuyingOptionDistribution `json:"buyingOptionDistributions"`
	CategoryDistributions     []CategoryDistribution     `json:"categoryDistributions"`
	ConditionDistributions    []ConditionDistribution    `json:"conditionDistributions"`
	DominantCategoryId        string                     `json:"dominantCategoryId"`
}

type AspectDistribution struct {
	AspectValueDistributions []AspectValueDistribution `json:"aspectValueDistributions"`
	LocalizedAspectName      string                    `json:"localizedAspectName"`
}

type AspectValueDistribution struct {
	LocalizedAspectValue string `json:"localizedAspectValue"`
	MatchCount           int    `json:"matchCount"`
	RefinementHref       string `json:"refinementHref"`
}

type BuyingOptionDistribution struct {
	BuyingOption   string `json:"buyingOption"`
	MatchCount     int    `json:"matchCount"`
	RefinementHref string `json:"refinementHref"`
}

type CategoryDistribution struct {
	CategoryId     string `json:"categoryId"`
	CategoryName   string `json:"categoryName"`
	MatchCount     int    `json:"matchCount"`
	RefinementHref string `json:"refinementHref"`
}

type ConditionDistribution struct {
	Condition      string `json:"condition"`
	ConditionId    string `json:"conditionId"`
	MatchCount     int    `json:"matchCount"`
	RefinementHref string `json:"refinementHref"`
}

type ErrorDetailV3 struct {
	Category     string             `json:"category"`
	Domain       string             `json:"domain"`
	ErrorId      int                `json:"errorId"`
	InputRefIds  []string           `json:"inputRefIds"`
	LongMessage  string             `json:"longMessage"`
	Message      string             `json:"message"`
	OutputRefIds []string           `json:"outputRefIds"`
	Parameters   []ErrorParameterV3 `json:"parameters"`
	Subdomain    string             `json:"subdomain"`
}

type ErrorParameterV3 struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
