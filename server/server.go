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

	s := r.Group("/shifts")
	{
		ctrl := controller.ShiftController{}
		// Shift
		// s.GET("", ctrl.IndexShift)
		s.POST("", ctrl.CreateShift)
		s.GET("", ctrl.ShowShift)
		s.GET("/:id	", ctrl.ShowShiftByUser)
		s.DELETE("/:id", ctrl.DeleteShift)

		// Shift Request
		// s.GET("/requests", ctrl.IndexRequest)
		s.POST("/requests", ctrl.CreateRequest)
		s.GET("/requests/:id", ctrl.ShowRequest)
		s.DELETE("/requests/:id", ctrl.DeleteRequest)

		// Shift Schedule
		s.POST("/schedule", ctrl.CreateSchedule)
		s.GET("schedule/:id", ctrl.ShowSchedule)
	}

	sr := r.Group("/stores")
	{
		ctrl := controller.StoreController{}
		sr.GET("/:id", ctrl.ShowStore)
		sr.POST("", ctrl.CreateStore)
	}

	wl := r.Group("/logs")
	{
		ctrl := controller.WorkLogController{}
		wl.POST("", ctrl.CreateUserLog)
		wl.PUT("/:id", ctrl.UpdateUserLogs)
		wl.GET("/:id", ctrl.ShowUserLogs)
	}

	return r
}
