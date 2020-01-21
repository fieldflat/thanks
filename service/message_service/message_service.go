package message

import (
	"github.com/gin-gonic/gin"

	"../../db"
	"../../entity"
)

// Service procides message's behavior
type Service struct{}

// Message is alias of entity.Message struct
type Message entity.Message

// GetAll is get all Message
func (s Service) GetAll() ([]Message, error) {
	db := db.GetDB()
	var m []Message

	if err := db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// CreateModel is create Message model
func (s Service) CreateModel(c *gin.Context) (Message, error) {
	db := db.GetDB()
	var m Message

	if err := c.BindJSON(&m); err != nil {
		return m, err
	}

	if err := db.Create(&m).Error; err != nil {
		return m, err
	}

	return m, nil
}

// GetByID is get a Message
func (s Service) GetByID(id string) (Message, error) {
	db := db.GetDB()
	var m Message

	if err := db.Where("id = ?", id).First(&m).Error; err != nil {
		return m, err
	}

	return m, nil
}

// UpdateByID is update a User
func (s Service) UpdateByID(id string, c *gin.Context) (Message, error) {
	db := db.GetDB()
	var m Message

	if err := db.Where("id = ?", id).First(&m).Error; err != nil {
		return m, err
	}

	if err := c.BindJSON(&m); err != nil {
		return m, err
	}

	db.Save(&m)

	return m, nil
}

// DeleteByID is delete a Message
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var m Message

	if err := db.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
