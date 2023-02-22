package crm_structures

type GetAccounts struct {
	Users []struct {
		ID          int    `json:"id"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	} `json:"users"`
	SalesAccount struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Address           interface{} `json:"address"`
		City              interface{} `json:"city"`
		State             interface{} `json:"state"`
		Zipcode           interface{} `json:"zipcode"`
		Country           interface{} `json:"country"`
		NumberOfEmployees interface{} `json:"number_of_employees"`
		AnnualRevenue     interface{} `json:"annual_revenue"`
		Website           interface{} `json:"website"`
		Phone             interface{} `json:"phone"`
		OpenDealsAmount   string      `json:"open_deals_amount"`
		LastContacted     interface{} `json:"last_contacted"`
		LastContactedMode interface{} `json:"last_contacted_mode"`
		Facebook          interface{} `json:"facebook"`
		Twitter           interface{} `json:"twitter"`
		Linkedin          interface{} `json:"linkedin"`
		Links             struct {
			Conversations string `json:"conversations"`
		} `json:"links"`
		CustomField struct {
			CfDomainName string `json:"cf_domain_name"`
		} `json:"custom_field"`
		UpdatedAt      string      `json:"updated_at"`
		OpenDealsCount int         `json:"open_deals_count"`
		Avatar         interface{} `json:"avatar"`
		OwnerID        int         `json:"owner_id"`
	} `json:"sales_account"`
}
type CreateAccounts struct {
	SalesAccount struct {
		Name        string `json:"name"`
		CustomField struct {
			CfDomainName string `json:"cf_domain_name"`
		} `json:"custom_field"`
	} `json:"sales_account"`
}
type UpdateAccounts struct {
	SalesAccount struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Address           interface{} `json:"address"`
		City              interface{} `json:"city"`
		State             interface{} `json:"state"`
		Zipcode           interface{} `json:"zipcode"`
		Country           interface{} `json:"country"`
		NumberOfEmployees interface{} `json:"number_of_employees"`
		AnnualRevenue     interface{} `json:"annual_revenue"`
		Website           interface{} `json:"website"`
		Phone             interface{} `json:"phone"`
		OpenDealsAmount   string      `json:"open_deals_amount"`
		LastContacted     interface{} `json:"last_contacted"`
		LastContactedMode interface{} `json:"last_contacted_mode"`
		Facebook          interface{} `json:"facebook"`
		Twitter           interface{} `json:"twitter"`
		Linkedin          interface{} `json:"linkedin"`
		Links             struct {
			Conversations string `json:"conversations"`
		} `json:"links"`
		CustomField struct {
			CfDomainName string `json:"cf_domain_name"`
		} `json:"custom_field"`
		UpdatedAt      string      `json:"updated_at"`
		OpenDealsCount int         `json:"open_deals_count"`
		Avatar         interface{} `json:"avatar"`
	} `json:"sales_account"`
}
type FieldsAccounts struct {
	Fields []struct {
		ID         int           `json:"id"`
		Label      string        `json:"label"`
		Name       string        `json:"name"`
		Type       string        `json:"type"`
		Default    bool          `json:"default"`
		Actionable bool          `json:"actionable"`
		Position   int           `json:"position"`
		Choices    []interface{} `json:"choices"`
		BaseModel  string        `json:"base_model"`
		Required   bool          `json:"required"`
	} `json:"fields"`
}
