package entity

// User is user models property
type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Sex          string `json:"sex"`
	Introduction string `json:"introduction"`
}
