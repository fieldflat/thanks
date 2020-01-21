package server

import (
	message "../controller/message_controller"
	user "../controller/user_controller"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := user.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	m := r.Group("/messages")
	{
		ctrl := message.Controller{}
		m.GET("", ctrl.Index)
		m.GET("/:id", ctrl.Show)
		m.POST("", ctrl.Create)
		m.PUT("/:id", ctrl.Update)
		m.DELETE("/:id", ctrl.Delete)
	}

	return r
}
