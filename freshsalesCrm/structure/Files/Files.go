package crm_structures

import "time"

type CreateFiles struct {
	URL            string `json:"url"`
	IsShared       bool   `json:"is_shared"`
	TargetableID   int    `json:"targetable_id"`
	TargetableType string `json:"targetable_type"`
	Name           string `json:"name"`
}
type GetFiles struct {
	Users []struct {
		ID           int         `json:"id"`
		DisplayName  string      `json:"display_name"`
		Email        string      `json:"email"`
		IsActive     bool        `json:"is_active"`
		WorkNumber   string      `json:"work_number"`
		MobileNumber interface{} `json:"mobile_number"`
	} `json:"users"`
	Documents []struct {
		ID                      int       `json:"id"`
		URL                     string    `json:"url"`
		Name                    string    `json:"name"`
		IsShared                bool      `json:"is_shared"`
		ContentFileSize         int       `json:"content_file_size"`
		ContentContentType      string    `json:"content_content_type"`
		ContentUpdatedAt        time.Time `json:"content_updated_at"`
		ContentFileSizeReadable string    `json:"content_file_size_readable"`
		CreatedAt               time.Time `json:"created_at"`
		IsAttached              bool      `json:"is_attached"`
		CreaterID               int       `json:"creater_id"`
	} `json:"documents"`
	DocumentLinks []struct {
		ID         int       `json:"id"`
		URL        string    `json:"url"`
		Name       string    `json:"name"`
		IsShared   bool      `json:"is_shared"`
		CreatedAt  time.Time `json:"created_at"`
		IsAttached bool      `json:"is_attached"`
		CreaterID  int       `json:"creater_id"`
	} `json:"document_links"`
}
