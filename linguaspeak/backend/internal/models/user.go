package models

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	GoogleID      string `json:"google_id,omitempty"`
	Name          string `json:"name"`
	ProfilePicURL string `json:"profile_pic_url,omitempty"`
}
