package crm_structures

import "time"

type CreateNotes struct {
	Note struct {
		Description    string `json:"description"`
		TargetableType string `json:"targetable_type"`
		TargetableID   int    `json:"targetable_id"`
	} `json:"note"`
}
type GetNotes struct {
	Note struct {
		ID          int       `json:"id"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"note"`
}
type UpdateNotes struct {
	Note struct {
		Description    string `json:"description"`
		TargetableType string `json:"targetable_type"`
		TargetableID   int    `json:"targetable_id"`
	} `json:"note"`
}
