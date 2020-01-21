package entity

// Message is Message models property
type Message struct {
	ID      uint   `json:"id"`
	Content string `json:"content" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
}
