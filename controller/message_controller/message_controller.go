package message

import (
	"fmt"

	"github.com/gin-gonic/gin"

	message "../../service/message_service"
)

// Controller is message controlller
type Controller struct{}

// Index action: GET /messages
func (pc Controller) Index(c *gin.Context) {
	var s message.Service
	userID, err1 := c.GetQuery("user_id")

	// クエリパラメータの指定がない場合
	if !err1 {
		p, err := s.GetAll()
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, p)
		}
	} else { // クエリパラメータの指定がある場合
		p, err2 := s.GetByUserID(userID)

		if err2 != nil {
			c.AbortWithStatus(404)
			fmt.Println(err2)
		} else {
			c.JSON(200, p)
		}
	}
}

// GetMessagesByUserID action: GET /messages?user_id=***
func (pc Controller) GetMessagesByUserID(c *gin.Context) {
	var s message.Service
	userID, err1 := c.GetQuery("user_id")

	if !err1 {
		c.AbortWithStatus(404)
		fmt.Println(err1)
	}

	p, err2 := s.GetByUserID(userID)

	if err2 != nil {
		c.AbortWithStatus(404)
		fmt.Println(err2)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /messages
func (pc Controller) Create(c *gin.Context) {
	var s message.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /messages/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s message.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /messages/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s message.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /messages/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s message.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
