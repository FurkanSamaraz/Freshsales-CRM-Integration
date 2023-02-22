package crm_structures

import "time"

type Create_Appointments struct {
	Appointment struct {
		Title                          string `json:"title"`
		Description                    string `json:"description"`
		FromDate                       string `json:"from_date"`
		EndDate                        string `json:"end_date"`
		TimeZone                       string `json:"time_zone"`
		Location                       string `json:"location"`
		TargetableID                   string `json:"targetable_id"`
		TargetableType                 string `json:"targetable_type"`
		AppointmentAttendeesAttributes []struct {
			AttendeeType string `json:"attendee_type"`
			AttendeeID   string `json:"attendee_id"`
		} `json:"appointment_attendees_attributes"`
	} `json:"appointment"`
}
type Get_Appointments struct {
	AppointmentAttendees []struct {
		ID       int `json:"id"`
		Attendee struct {
			Type string `json:"type"`
			ID   int    `json:"id"`
		} `json:"attendee"`
	} `json:"appointment_attendees"`
	Users []struct {
		ID          int         `json:"id"`
		DisplayName string      `json:"display_name"`
		Avatar      interface{} `json:"avatar"`
		Type        string      `json:"type"`
	} `json:"users"`
	Leads []struct {
		ID          int         `json:"id"`
		DisplayName string      `json:"display_name"`
		Avatar      interface{} `json:"avatar"`
		Type        string      `json:"type,omitempty"`
	} `json:"leads"`
	Appointment struct {
		ID                     int       `json:"id"`
		TimeZone               string    `json:"time_zone"`
		Title                  string    `json:"title"`
		Description            string    `json:"description"`
		Location               string    `json:"location"`
		FromDate               time.Time `json:"from_date"`
		EndDate                time.Time `json:"end_date"`
		Latitude               string    `json:"latitude"`
		Longitude              string    `json:"longitude"`
		CheckedinAt            string    `json:"checkedin_at"`
		AppointmentAttendeeIds []int     `json:"appointment_attendee_ids"`
		Targetable             struct {
			Type string `json:"type"`
			ID   int    `json:"id"`
		} `json:"targetable"`
		CreaterID int `json:"creater_id"`
	} `json:"appointment"`
}
type Update_Appointments struct {
	Appointment struct {
		Title                          string `json:"title"`
		Description                    string `json:"description"`
		FromDate                       string `json:"from_date"`
		EndDate                        string `json:"end_date"`
		TimeZone                       string `json:"time_zone"`
		Location                       string `json:"location"`
		TargetableID                   string `json:"targetable_id"`
		TargetableType                 string `json:"targetable_type"`
		AppointmentAttendeesAttributes []struct {
			AttendeeType string `json:"attendee_type"`
			AttendeeID   string `json:"attendee_id"`
		} `json:"appointment_attendees_attributes"`
	} `json:"appointment"`
}
