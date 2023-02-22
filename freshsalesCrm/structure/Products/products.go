package crm_structures

import "time"

type CreateProducts struct {
	Product struct {
		Name          string    `json:"name"`
		Description   string    `json:"description"`
		Category      string    `json:"category"`
		IsActive      bool      `json:"is_active"`
		ParentProduct int       `json:"parent_product"`
		ProductCode   string    `json:"product_code"`
		SkuNumber     string    `json:"sku_number"`
		ValidTill     time.Time `json:"valid_till"`
		OwnerID       int       `json:"owner_id"`
		CustomField   struct {
			CfRating   string `json:"cf_rating"`
			CfQuantity int    `json:"cf_quantity"`
		} `json:"custom_field"`
	} `json:"product"`
}
type GetProducts struct {
	Product struct {
		ID                 int           `json:"id"`
		Name               string        `json:"name"`
		Category           string        `json:"category"`
		ProductCode        string        `json:"product_code"`
		SkuNumber          string        `json:"sku_number"`
		ParentProduct      int           `json:"parent_product"`
		TerritoryID        interface{}   `json:"territory_id"`
		ValidTill          time.Time     `json:"valid_till"`
		IsActive           bool          `json:"is_active"`
		OwnerID            int           `json:"owner_id"`
		IsDeleted          bool          `json:"is_deleted"`
		CreatedAt          time.Time     `json:"created_at"`
		UpdatedAt          time.Time     `json:"updated_at"`
		PricingType        int           `json:"pricing_type"`
		ProductPricings    []interface{} `json:"product_pricings"`
		Avatar             interface{}   `json:"avatar"`
		BaseCurrencyAmount interface{}   `json:"base_currency_amount"`
		CreaterID          int           `json:"creater_id"`
		UpdaterID          interface{}   `json:"updater_id"`
		Description        string        `json:"description"`
		LookupInformation  struct {
		} `json:"lookup_information"`
	} `json:"product"`
}
type UpdateProducts struct {
	Product struct {
		Category    string `json:"category"`
		IsActive    bool   `json:"is_active"`
		CustomField struct {
			CfRating string `json:"cf_rating"`
		} `json:"custom_field"`
	} `json:"product"`
}
