package crm_structures

type GetSearch struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	CompanyName string `json:"company_name,omitempty"`
}
type CreateSearch struct {
	FilterRule []struct {
		Attribute string `json:"attribute"`
		Operator  string `json:"operator"`
		Value     string `json:"value"`
	} `json:"filter_rule"`
}
