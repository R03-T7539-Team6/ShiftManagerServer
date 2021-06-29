package server

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/controller"
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
		ctrl := controller.UserController{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	sr := r.Group("/stores")
	{
		ctrl := controller.StoreController{}
		sr.GET("/:id", ctrl.Show)
		sr.POST("", ctrl.Create)
	}

	return r
}
