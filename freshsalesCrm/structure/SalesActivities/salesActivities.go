package crm_structures

import "time"

type CreateSalesActivities struct {
	SalesActivity struct {
		Title                  string    `json:"title"`
		Notes                  string    `json:"notes"`
		TargetableID           string    `json:"targetable_id"`
		TargetableType         string    `json:"targetable_type"`
		StartDate              time.Time `json:"start_date"`
		EndDate                time.Time `json:"end_date"`
		OwnerID                string    `json:"owner_id"`
		SalesActivityTypeID    string    `json:"sales_activity_type_id"`
		SalesActivityOutcomeID string    `json:"sales_activity_outcome_id"`
		CustomField            struct {
			CfCustomerOption bool `json:"cf_customer_option"`
		} `json:"custom_field"`
	} `json:"sales_activity"`
}
type GetSalesActivities struct {
	SalesActivity struct {
		ID                     int         `json:"id"`
		Title                  string      `json:"title"`
		StartDate              time.Time   `json:"start_date"`
		EndDate                time.Time   `json:"end_date"`
		SalesActivityTypeID    int         `json:"sales_activity_type_id"`
		CreatedAt              time.Time   `json:"created_at"`
		UpdatedAt              time.Time   `json:"updated_at"`
		Notes                  interface{} `json:"notes"`
		OwnerID                int         `json:"owner_id"`
		SalesActivityOutcomeID interface{} `json:"sales_activity_outcome_id"`
		RemoteID               interface{} `json:"remote_id"`
		ImportID               interface{} `json:"import_id"`
		CreaterID              int         `json:"creater_id"`
		UpdaterID              interface{} `json:"updater_id"`
		ConversationTime       time.Time   `json:"conversation_time"`
		TargetableID           int         `json:"targetable_id"`
		TargetableType         string      `json:"targetable_type"`
		Latitude               string      `json:"latitude"`
		Longitude              string      `json:"longitude"`
		Location               string      `json:"location"`
		CheckedinAt            time.Time   `json:"checkedin_at"`
		CustomField            struct {
			CfCustomerOption bool `json:"cf_customer_option"`
		} `json:"custom_field"`
	} `json:"sales_activity"`
}
type UpdateSalesActivities struct {
	Title                  string `json:"title"`
	Notes                  string `json:"notes"`
	TargetableID           string `json:"targetable_id"`
	TargetableType         string `json:"targetable_type"`
	OwnerID                string `json:"owner_id"`
	SalesActivityTypeID    string `json:"sales_activity_type_id"`
	SalesActivityOutcomeID string `json:"sales_activity_outcome_id"`
}
