package controller

import (
	"fmt"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/gin-gonic/gin"
)

type ShiftController struct{}

// IndexRequest action: GET /shift
// func (sc ShiftController) IndexShift(c *gin.Context) {
// 	var sr model.Shift
// 	p, err := sr.GetAll()
// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, p)
// 	}
// }

// CreateShift action: POST /shift
func (sc ShiftController) CreateShift(c *gin.Context) {
	var sr model.Shift
	p, err := sr.CreateShift(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// ShowShift action: GET /shift/:id
func (sc ShiftController) ShowShift(c *gin.Context) {
	id := c.Params.ByName("id")
	var sr model.Shift
	p, err := sr.GetByUserId(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// DeleteShift action: DELETE /shift/:id
func (sc ShiftController) DeleteShift(c *gin.Context) {
	id := c.Params.ByName("id")
	var sr model.Shift
	if err := sr.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + id: "deleted",
		})
	}
}

// *******************ShiftRequest
// IndexRequest action: GET /shift/requests
// func (sc ShiftController) IndexRequest(c *gin.Context) {
// 	var sr model.ShiftRequest
// 	p, err := sr.GetAll()
// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, p)
// 	}
// }

// CreateRequest action: POST /shift/requests
func (sc ShiftController) CreateRequest(c *gin.Context) {
	var sr model.ShiftRequest
	p, err := sr.CreateShiftRequest(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// ShowRequest action: GET /shift/requests/:id
func (sc ShiftController) ShowRequest(c *gin.Context) {
	id := c.Params.ByName("id")
	var sr model.ShiftRequest
	p, err := sr.GetByUserId(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// DeleteRequest action: DELETE /shift/requests/:id
func (sc ShiftController) DeleteRequest(c *gin.Context) {
	id := c.Params.ByName("id")
	var sr model.ShiftRequest
	if err := sr.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + id: "deleted",
		})
	}
}

// ****************ShiftSchedule
// CreateSchedule action: POST /shift/schedule/
func (sc ShiftController) CreateSchedule(c *gin.Context) {
	var ss model.ShiftSchedule
	p, err := ss.CreateShiftSchedule(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// ShowRequest action: GET /shift/schedule/:id
func (sc ShiftController) ShowSchedule(c *gin.Context) {
	id := c.Params.ByName("id")
	var ss model.ShiftSchedule
	p, err := ss.GetById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
