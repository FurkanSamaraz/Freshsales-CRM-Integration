package crm_structures

import "time"

type GetJobStatus struct {
	Status          string    `json:"status"`
	JobTypeName     string    `json:"job_type_name"`
	CreatedAt       time.Time `json:"created_at"`
	StatusUpdatedAt time.Time `json:"status_updated_at"`
	Progress        int       `json:"progress"`
	Data            struct {
		RecordStatus struct {
			Succeeded int `json:"succeeded"`
			Failed    int `json:"failed"`
		} `json:"record_status"`
		DetailedFailureReport []struct {
			Emails       string   `json:"emails,omitempty"`
			ErrorMessage []string `json:"error_message"`
			MobileNumber string   `json:"mobile_number,omitempty"`
		} `json:"detailed_failure_report"`
	} `json:"data"`
}
