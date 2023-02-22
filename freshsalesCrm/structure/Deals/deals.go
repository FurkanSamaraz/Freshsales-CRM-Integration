package crm_structures

type CreateDeals struct {
	Deal struct {
		Name           string `json:"name"`
		Amount         int    `json:"amount"`
		SalesAccountID int    `json:"sales_account_id"`
		CustomField    struct {
			CfNumberOfAgents int `json:"cf_number_of_agents"`
		} `json:"custom_field"`
	} `json:"deal"`
}
type GetDeals struct {
	Users []struct {
		ID          int    `json:"id"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	} `json:"users"`
	Deal struct {
		ID                 int         `json:"id"`
		Name               string      `json:"name"`
		Amount             string      `json:"amount"`
		BaseCurrencyAmount string      `json:"base_currency_amount"`
		ExpectedClose      interface{} `json:"expected_close"`
		ClosedDate         interface{} `json:"closed_date"`
		StageUpdatedTime   string      `json:"stage_updated_time"`
		CustomField        struct {
			CfNumberOfAgents int `json:"cf_number_of_agents"`
		} `json:"custom_field"`
		Probability interface{} `json:"probability"`
		UpdatedAt   string      `json:"updated_at"`
		CreatedAt   string      `json:"created_at"`
		Age         interface{} `json:"age"`
		OwnerID     int         `json:"owner_id"`
	} `json:"deal"`
}
type UpdateDeals struct {
	UniqueIdentifier struct {
		Name string `json:"name"`
	} `json:"unique_identifier"`
	Deal struct {
		Amount string `json:"amount"`
	} `json:"deal"`
}
