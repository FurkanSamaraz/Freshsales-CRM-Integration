package crm_structures

import "time"

type CreateContacts struct {
	Contact struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		MobileNumber string `json:"mobile_number"`
		CustomField  struct {
			CfIsActive bool `json:"cf_is_active"`
		} `json:"custom_field"`
		SalesAccounts []struct {
			ID        int  `json:"id"`
			IsPrimary bool `json:"is_primary"`
		} `json:"sales_accounts"`
	} `json:"contact"`
}
type GetContacts struct {
	Contact struct {
		ID              int         `json:"id"`
		FirstName       string      `json:"first_name"`
		LastName        string      `json:"last_name"`
		DisplayName     string      `json:"display_name"`
		Avatar          interface{} `json:"avatar"`
		JobTitle        interface{} `json:"job_title"`
		City            interface{} `json:"city"`
		State           interface{} `json:"state"`
		Zipcode         interface{} `json:"zipcode"`
		Country         interface{} `json:"country"`
		Email           interface{} `json:"email"`
		TimeZone        interface{} `json:"time_zone"`
		WorkNumber      interface{} `json:"work_number"`
		MobileNumber    string      `json:"mobile_number"`
		Address         interface{} `json:"address"`
		LastSeen        interface{} `json:"last_seen"`
		LeadScore       int         `json:"lead_score"`
		LastContacted   interface{} `json:"last_contacted"`
		OpenDealsAmount string      `json:"open_deals_amount"`
		Links           struct {
			Conversations string `json:"conversations"`
			Activities    string `json:"activities"`
		} `json:"links"`
		CustomField struct {
		} `json:"custom_field"`
		UpdatedAt string      `json:"updated_at"`
		Keyword   interface{} `json:"keyword"`
		Medium    interface{} `json:"medium"`
		Facebook  interface{} `json:"facebook"`
		Twitter   interface{} `json:"twitter"`
		Linkedin  interface{} `json:"linkedin"`
	} `json:"contact"`
}
type UpdateContacts struct {
	Contact struct {
		MobileNumber string `json:"mobile_number"`
		CustomField  struct {
			CfIsActive bool `json:"cf_is_active"`
		} `json:"custom_field"`
	} `json:"contact"`
}
type FieldsContacts struct {
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
type UpsertContacts struct {
	UniqueIdentifier struct {
		Emails string `json:"emails"`
	} `json:"unique_identifier"`
	Contact struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		MobileNumber string `json:"mobile_number"`
	} `json:"contact"`
}
type ActivitieContacts struct {
	Users []struct {
		ID          int    `json:"id"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	} `json:"users"`
	Activities []struct {
		CompositeID string `json:"composite_id"`
		ActionData  struct {
			StageID   int    `json:"stage_id"`
			StageName string `json:"stage_name"`
		} `json:"action_data,omitempty"`
		ActionType     string      `json:"action_type"`
		CreatedAt      string      `json:"created_at"`
		UserActivity   bool        `json:"user_activity"`
		TargetableType string      `json:"targetable_type"`
		TargetableID   int         `json:"targetable_id"`
		ActionableType interface{} `json:"actionable_type"`
		ActionableID   interface{} `json:"actionable_id"`
		TargetableName string      `json:"targetable_name"`
		UserID         int         `json:"user_id"`
		ActionData0    struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
		} `json:"action_data,omitempty"`
		ActionData1 struct {
			Description string    `json:"description"`
			EndDate     time.Time `json:"end_date"`
			FromDate    time.Time `json:"from_date"`
			TimeZone    string    `json:"time_zone"`
			Title       string    `json:"title"`
		} `json:"action_data,omitempty"`
	} `json:"activities"`
}
