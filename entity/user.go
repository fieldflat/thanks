package entity

// User is user models property
type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Sex          string `json:"sex"`
	Introduction string `json:"introduction"`
}
