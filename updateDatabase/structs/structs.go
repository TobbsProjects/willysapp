package structs

type WillysJson struct {
	Results []struct {
		PotentialPromotions []struct {
			Applied               bool `json:"applied"`
			LowestHistoricalPrice struct {
				CurrencyIso    string      `json:"currencyIso"`
				Value          float64     `json:"value"`
				PriceType      string      `json:"priceType"`
				FormattedValue string      `json:"formattedValue"`
				MinQuantity    interface{} `json:"minQuantity"`
				MaxQuantity    interface{} `json:"maxQuantity"`
			} `json:"lowestHistoricalPrice"`
			CampaignType   string `json:"campaignType"`
			PromotionTheme struct {
				Code    string      `json:"code"`
				Visible interface{} `json:"visible"`
			} `json:"promotionTheme"`
			TextLabelGenerated      string      `json:"textLabelGenerated"`
			TextLabelManual         interface{} `json:"textLabelManual"`
			ConditionLabelFormatted string      `json:"conditionLabelFormatted"`
			PromotionRedeemLimit    interface{} `json:"promotionRedeemLimit"`
			PromotionPercentage     interface{} `json:"promotionPercentage"`
			CartLabelFormatted      string      `json:"cartLabelFormatted"`
			Price                   struct {
				CurrencyIso    string      `json:"currencyIso"`
				Value          float64     `json:"value"`
				PriceType      string      `json:"priceType"`
				FormattedValue string      `json:"formattedValue"`
				MinQuantity    interface{} `json:"minQuantity"`
				MaxQuantity    interface{} `json:"maxQuantity"`
			} `json:"price"`
			ProductCodes     interface{} `json:"productCodes"`
			PromotionType    string      `json:"promotionType"`
			ComparePrice     string      `json:"comparePrice"`
			ConditionLabel   string      `json:"conditionLabel"`
			RewardLabel      string      `json:"rewardLabel"`
			TextLabel        string      `json:"textLabel"`
			RedeemLimitLabel string      `json:"redeemLimitLabel"`
			CartLabel        interface{} `json:"cartLabel"`
			ValidUntil       interface{} `json:"validUntil"`
			TimesUsed        interface{} `json:"timesUsed"`
			QualifyingCount  interface{} `json:"qualifyingCount"`
			SplashTitleText  string      `json:"splashTitleText"`
			RealMixAndMatch  bool        `json:"realMixAndMatch"`
			MainProductCode  string      `json:"mainProductCode"`
			Threshold        float64     `json:"threshold"`
			Code             string      `json:"code"`
			Priority         int         `json:"priority"`
		} `json:"potentialPromotions"`
		PriceValue              float64     `json:"priceValue"`
		ProductLine2            string      `json:"productLine2"`
		PickupProductLine2      string      `json:"pickupProductLine2"`
		PriceUnit               string      `json:"priceUnit"`
		PriceNoUnit             string      `json:"priceNoUnit"`
		SavingsAmount           float64     `json:"savingsAmount"`
		GoogleAnalyticsCategory string      `json:"googleAnalyticsCategory"`
		GdprTrackingIncompliant bool        `json:"gdprTrackingIncompliant"`
		Price                   string      `json:"price"`
		DepositPrice            string      `json:"depositPrice"`
		AverageWeight           float64     `json:"averageWeight"`
		ComparePrice            string      `json:"comparePrice"`
		ComparePriceUnit        string      `json:"comparePriceUnit"`
		IsDrugProduct           bool        `json:"isDrugProduct"`
		SolrSearchScore         float64     `json:"solrSearchScore"`
		EnergyDeclaration       interface{} `json:"energyDeclaration"`
		Image                   struct {
			ImageType    string      `json:"imageType"`
			Format       string      `json:"format"`
			URL          string      `json:"url"`
			AltText      interface{} `json:"altText"`
			GalleryIndex interface{} `json:"galleryIndex"`
			Width        interface{} `json:"width"`
		} `json:"image"`
		Ranking                    float64  `json:"ranking"`
		NewsSplashProduct          bool     `json:"newsSplashProduct"`
		NotAllowedB2B              bool     `json:"notAllowedB2b"`
		NotAllowedAnonymous        bool     `json:"notAllowedAnonymous"`
		ShowGoodPriceIcon          bool     `json:"showGoodPriceIcon"`
		NicotineMedicalProduct     bool     `json:"nicotineMedicalProduct"`
		TobaccoFreeNicotineProduct bool     `json:"tobaccoFreeNicotineProduct"`
		NonEkoProduct              bool     `json:"nonEkoProduct"`
		TobaccoProduct             bool     `json:"tobaccoProduct"`
		Labels                     []string `json:"labels"`
		Online                     bool     `json:"online"`
		DisplayVolume              string   `json:"displayVolume"`
		Manufacturer               string   `json:"manufacturer"`
		IncrementValue             float64  `json:"incrementValue"`
		ProductBasketType          struct {
			Code string `json:"code"`
			Type string `json:"type"`
		} `json:"productBasketType"`
		Thumbnail struct {
			ImageType    string      `json:"imageType"`
			Format       string      `json:"format"`
			URL          string      `json:"url"`
			AltText      interface{} `json:"altText"`
			GalleryIndex interface{} `json:"galleryIndex"`
			Width        interface{} `json:"width"`
		} `json:"thumbnail"`
		OutOfStock        bool   `json:"outOfStock"`
		AddToCartDisabled bool   `json:"addToCartDisabled"`
		Code              string `json:"code"`
		Name              string `json:"name"`
	} `json:"results"`
	Sorts []struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		Selected bool   `json:"selected"`
	} `json:"sorts"`
	Pagination struct {
		PageSize                     int    `json:"pageSize"`
		CurrentPage                  int    `json:"currentPage"`
		Sort                         string `json:"sort"`
		NumberOfPages                int    `json:"numberOfPages"`
		TotalNumberOfResults         int    `json:"totalNumberOfResults"`
		AllProductsInCategoriesCount int    `json:"allProductsInCategoriesCount"`
		AllProductsInSearchCount     int    `json:"allProductsInSearchCount"`
	} `json:"pagination"`
	RelatedResults           interface{}   `json:"relatedResults"`
	RelatedResultsPagination interface{}   `json:"relatedResultsPagination"`
	CurrentQuery             interface{}   `json:"currentQuery"`
	Breadcrumbs              []interface{} `json:"breadcrumbs"`
	Facets                   []struct {
		Code        string      `json:"code"`
		Name        string      `json:"name"`
		Priority    int         `json:"priority"`
		Category    bool        `json:"category"`
		MultiSelect bool        `json:"multiSelect"`
		Visible     bool        `json:"visible"`
		TopValues   interface{} `json:"topValues"`
		Values      []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Count int    `json:"count"`
			Query struct {
				URL   string `json:"url"`
				Query struct {
					Value              string        `json:"value"`
					FilterQueries      []interface{} `json:"filterQueries"`
					SearchQueryContext interface{}   `json:"searchQueryContext"`
					SearchType         interface{}   `json:"searchType"`
				} `json:"query"`
			} `json:"query"`
			Selected bool `json:"selected"`
		} `json:"values"`
	} `json:"facets"`
	FreeTextSearch      string      `json:"freeTextSearch"`
	CategoryCode        interface{} `json:"categoryCode"`
	KeywordRedirectURL  interface{} `json:"keywordRedirectUrl"`
	SpellingSuggestion  interface{} `json:"spellingSuggestion"`
	CategoryName        interface{} `json:"categoryName"`
	CustomSuggestion    interface{} `json:"customSuggestion"`
	CategoryBreadcrumbs struct {
		URL          string      `json:"url"`
		Name         string      `json:"name"`
		LinkClass    interface{} `json:"linkClass"`
		CategoryCode string      `json:"categoryCode"`
	} `json:"categoryBreadcrumbs"`
	NavigationCategories interface{}   `json:"navigationCategories"`
	SubCategories        []interface{} `json:"subCategories"`
	CategoryInfo         struct {
		Code               string      `json:"code"`
		Name               string      `json:"name"`
		URL                string      `json:"url"`
		Description        interface{} `json:"description"`
		Image              interface{} `json:"image"`
		ParentCategoryName string      `json:"parentCategoryName"`
		Sequence           int         `json:"sequence"`
	} `json:"categoryInfo"`
	SuperCategories []struct {
		Code               string      `json:"code"`
		Name               string      `json:"name"`
		URL                string      `json:"url"`
		Description        interface{} `json:"description"`
		Image              interface{} `json:"image"`
		ParentCategoryName interface{} `json:"parentCategoryName"`
		Sequence           int         `json:"sequence"`
	} `json:"superCategories"`
	RestrictedProductTypes string `json:"restrictedProductTypes"`
}
