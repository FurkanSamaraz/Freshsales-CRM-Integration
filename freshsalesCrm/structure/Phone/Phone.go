package crm_structures

type CreatePhone struct {
	PhoneCall struct {
		CallDirection  bool   `json:"call_direction"`
		TargetableType string `json:"targetable_type"`
		Targetable     struct {
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			WorkNumber   string `json:"work_number"`
			MobileNumber string `json:"mobile_number"`
			Company      struct {
				Name string `json:"name"`
			} `json:"company"`
		} `json:"targetable"`
		Note struct {
			Description string `json:"description"`
		} `json:"note"`
	} `json:"phone_call"`
}
