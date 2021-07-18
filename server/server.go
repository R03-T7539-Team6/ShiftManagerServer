package server

import (
	"net/http"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/controller"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

/*************************************************
 *	specification;
 *	name 			= router
 *	Function 	= define routing endpoint
 *	note			= use Init function
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= None
 *  output    = *gin.Engine
 *  end of specification;
**************************************************/
func router() *gin.Engine {
	r := gin.Default()

	l := r.Group("/")
	{
		ctrl := controller.AuthorizationController{}
		l.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello ShiftManager",
			})
		})
		l.POST("/login", ctrl.Login)
		l.POST("/signup", ctrl.Signup)
	}
	u := r.Group("/users")
	{
		ctrl := controller.UserController{}
		// u.GET("", ctrl.Index)
		u.GET("", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	s := r.Group("/shifts")
	{
		ctrl := controller.ShiftController{}
		// Shift
		s.POST("", ctrl.CreateShift)
		s.GET("", ctrl.ShowShiftByUser)
		s.PUT("/:id", ctrl.UpdateShift)
		s.DELETE("/:id", ctrl.DeleteShift)

		// Shift Request
		s.POST("/requests", ctrl.CreateRequest)
		s.GET("/requests", ctrl.ShowRequest)
		s.DELETE("/requests/:id", ctrl.DeleteRequest)

		// Shift Schedule
		s.POST("/schedule", ctrl.CreateSchedule)
		s.GET("/schedule/:id", ctrl.ShowSchedule)
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
		wl.GET("", ctrl.ShowUserLogs)
	}

	return r
}
