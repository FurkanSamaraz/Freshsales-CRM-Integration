package crm_structures

import "time"

type CreateTasks struct {
	Task struct {
		Title               string `json:"title"`
		Description         string `json:"description"`
		DueDate             string `json:"due_date"`
		OwnerID             string `json:"owner_id"`
		TargetableID        string `json:"targetable_id"`
		TargetableType      string `json:"targetable_type"`
		TaskUsersAttributes []struct {
			UserID string `json:"user_id"`
		} `json:"task_users_attributes"`
	} `json:"task"`
}
type GetTasks struct {
	Users []struct {
		ID          int         `json:"id"`
		DisplayName string      `json:"display_name"`
		Avatar      interface{} `json:"avatar"`
		Type        string      `json:"type"`
	} `json:"users"`
	Leads []struct {
		ID          int         `json:"id"`
		Avatar      interface{} `json:"avatar"`
		DisplayName string      `json:"display_name"`
	} `json:"leads"`
	Task struct {
		ID          int       `json:"id"`
		Status      int       `json:"status"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		OwnerID     int       `json:"owner_id"`
		DueDate     time.Time `json:"due_date"`
		CreaterID   int       `json:"creater_id"`
		UserIds     []int     `json:"user_ids"`
		Targetable  struct {
			Type string `json:"type"`
			ID   int    `json:"id"`
		} `json:"targetable"`
	} `json:"task"`
}
type UpdateTasks struct {
	Task struct {
		Title               string `json:"title"`
		Description         string `json:"description"`
		DueDate             string `json:"due_date"`
		OwnerID             string `json:"owner_id"`
		TargetableID        string `json:"targetable_id"`
		TargetableType      string `json:"targetable_type"`
		TaskUsersAttributes []struct {
			UserID string `json:"user_id"`
		} `json:"task_users_attributes"`
	} `json:"task"`
}
