package crm_structures

type CreateLeads struct {
	Lead struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		MobileNumber string `json:"mobile_number"`
		Company      struct {
			Name string `json:"name"`
		} `json:"company"`
	} `json:"lead"`
}

type GetLeads struct {
	Lead struct {
		ID               int         `json:"id"`
		JobTitle         string      `json:"job_title"`
		Email            string      `json:"email"`
		WorkNumber       string      `json:"work_number"`
		MobileNumber     string      `json:"mobile_number"`
		Address          string      `json:"address"`
		City             string      `json:"city"`
		State            string      `json:"state"`
		Zipcode          string      `json:"zipcode"`
		Country          string      `json:"country"`
		TimeZone         interface{} `json:"time_zone"`
		DisplayName      string      `json:"display_name"`
		Avatar           string      `json:"avatar"`
		Keyword          string      `json:"keyword"`
		Medium           string      `json:"medium"`
		LastSeen         string      `json:"last_seen"`
		LastContacted    string      `json:"last_contacted"`
		LeadScore        int         `json:"lead_score"`
		StageUpdatedTime string      `json:"stage_updated_time"`
		FirstName        string      `json:"first_name"`
		LastName         string      `json:"last_name"`
		Company          struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Address           string      `json:"address"`
			City              string      `json:"city"`
			State             string      `json:"state"`
			Zipcode           string      `json:"zipcode"`
			Country           string      `json:"country"`
			NumberOfEmployees interface{} `json:"number_of_employees"`
			AnnualRevenue     interface{} `json:"annual_revenue"`
			Website           string      `json:"website"`
			Phone             string      `json:"phone"`
			IndustryTypeID    int         `json:"industry_type_id"`
			BusinessTypeID    int         `json:"business_type_id"`
		} `json:"company"`
		Deal  interface{} `json:"deal"`
		Links struct {
			Conversations string `json:"conversations"`
			Activities    string `json:"activities"`
		} `json:"links"`
		UpdatedAt string      `json:"updated_at"`
		Facebook  interface{} `json:"facebook"`
		Twitter   string      `json:"twitter"`
		Linkedin  string      `json:"linkedin"`
	} `json:"lead"`
}
type UpdateLeads struct {
	UniqueIdentifier struct {
		Emails string `json:"emails"`
	} `json:"unique_identifier"`
	Lead struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		MobileNumber string `json:"mobile_number"`
	} `json:"lead"`
}
type FieldsLeads struct {
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
type ActivitiesLeads struct {
	Users []struct {
		ID          int    `json:"id"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	} `json:"users"`
	Activities []struct {
		CompositeID    string      `json:"composite_id"`
		ActionData     string      `json:"action_data"`
		ActionType     string      `json:"action_type"`
		CreatedAt      string      `json:"created_at"`
		UserActivity   bool        `json:"user_activity"`
		TargetableType interface{} `json:"targetable_type"`
		TargetableID   interface{} `json:"targetable_id"`
		ActionableType interface{} `json:"actionable_type"`
		ActionableID   interface{} `json:"actionable_id"`
		TargetableName interface{} `json:"targetable_name"`
		UserID         interface{} `json:"user_id"`
	} `json:"activities"`
}
