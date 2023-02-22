package crm_structures

type CreateDocuments struct {
	CpqDocument struct {
		DealID                  int    `json:"deal_id"`
		SalesAccountID          int    `json:"sales_account_id"`
		ContactID               int    `json:"contact_id"`
		DocumentType            string `json:"document_type"`
		Stage                   string `json:"stage"`
		ValidTill               string `json:"valid_till"`
		ShippingAddress         string `json:"shipping_address"`
		ShippingCity            string `json:"shipping_city"`
		ShippingState           string `json:"shipping_state"`
		ShippingZipcode         string `json:"shipping_zipcode"`
		ShippingCountry         string `json:"shipping_country"`
		BillingAddress          string `json:"billing_address"`
		BillingCity             string `json:"billing_city"`
		BillingState            string `json:"billing_state"`
		BillingZipcode          string `json:"billing_zipcode"`
		BillingCountry          string `json:"billing_country"`
		Amount                  int    `json:"amount"`
		DisplayName             string `json:"display_name"`
		CurrencyCode            string `json:"currency_code"`
		OwnerID                 int    `json:"owner_id"`
		TerritoryID             int    `json:"territory_id"`
		CpqDocumentTemplateName string `json:"cpq_document_template_name"`
	} `json:"cpq_document"`
}
type GetDocuments struct {
	CpqDocument struct {
		ID                 int         `json:"id"`
		DealID             int         `json:"deal_id"`
		ContactID          int         `json:"contact_id"`
		SalesAccountID     int         `json:"sales_account_id"`
		DocumentType       string      `json:"document_type"`
		Stage              string      `json:"stage"`
		ValidTill          string      `json:"valid_till"`
		ShippingAddress    string      `json:"shipping_address"`
		ShippingCity       string      `json:"shipping_city"`
		ShippingState      string      `json:"shipping_state"`
		ShippingZipcode    string      `json:"shipping_zipcode"`
		ShippingCountry    string      `json:"shipping_country"`
		BillingAddress     string      `json:"billing_address"`
		BillingCity        string      `json:"billing_city"`
		BillingState       string      `json:"billing_state"`
		BillingZipcode     string      `json:"billing_zipcode"`
		BillingCountry     string      `json:"billing_country"`
		OwnerID            int         `json:"owner_id"`
		Amount             float64     `json:"amount"`
		CurrencyCode       string      `json:"currency_code"`
		BaseCurrencyAmount float64     `json:"base_currency_amount"`
		CreaterID          int         `json:"creater_id"`
		UpdaterID          interface{} `json:"updater_id"`
		CustomField        struct {
			CfCustomStatus interface{} `json:"cf_custom_status"`
			CfRating       interface{} `json:"cf_rating"`
		} `json:"custom_field"`
		CreatedAt           string        `json:"created_at"`
		UpdatedAt           string        `json:"updated_at"`
		IsDeleted           bool          `json:"is_deleted"`
		DocumentNumber      string        `json:"document_number"`
		TerritoryID         int           `json:"territory_id"`
		IsDealPrimary       bool          `json:"is_deal_primary"`
		Products            []interface{} `json:"products"`
		DisplayName         string        `json:"display_name"`
		HTMLContent         interface{}   `json:"html_content"`
		HTMLContentCSS      interface{}   `json:"html_content_css"`
		HTMLContentSettings struct {
		} `json:"html_content_settings"`
		HTMLHeader              interface{} `json:"html_header"`
		HTMLFooter              interface{} `json:"html_footer"`
		EmailTemplateID         interface{} `json:"email_template_id"`
		CpqDocumentTemplateName string      `json:"cpq_document_template_name"`
		HasProducts             bool        `json:"has_products"`
	} `json:"cpq_document"`
}
type UpdateDocuments struct {
	CpqDocument struct {
		DocumentType       string      `json:"document_type"`
		Stage              string      `json:"stage"`
		ValidTill          string      `json:"valid_till"`
		ShippingAddress    string      `json:"shipping_address"`
		ShippingCity       string      `json:"shipping_city"`
		ShippingState      string      `json:"shipping_state"`
		ShippingZipcode    string      `json:"shipping_zipcode"`
		ShippingCountry    string      `json:"shipping_country"`
		BillingAddress     string      `json:"billing_address"`
		BillingCity        string      `json:"billing_city"`
		BillingState       string      `json:"billing_state"`
		BillingZipcode     string      `json:"billing_zipcode"`
		BillingCountry     string      `json:"billing_country"`
		OwnerID            int         `json:"owner_id"`
		Amount             float64     `json:"amount"`
		CurrencyCode       string      `json:"currency_code"`
		BaseCurrencyAmount float64     `json:"base_currency_amount"`
		CreaterID          int         `json:"creater_id"`
		UpdaterID          interface{} `json:"updater_id"`
		CustomField        struct {
			CfCustomStatus interface{} `json:"cf_custom_status"`
			CfRating       interface{} `json:"cf_rating"`
		} `json:"custom_field"`
	} `json:"cpq_document"`
}
