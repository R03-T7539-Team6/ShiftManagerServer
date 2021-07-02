package controller

import (
	"fmt"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/gin-gonic/gin"
)

type WorkLogController struct{}

// CreateUserLog action: POST /logs
func (sc WorkLogController) CreateUserLog(c *gin.Context) {
	var s model.WorkLog
	p, err := s.CreateLog(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// ShowUserLogs action: GET /logs/:id
func (sc WorkLogController) ShowUserLogs(c *gin.Context) {
	id := c.Params.ByName("id")
	var s model.WorkLog
	p, err := s.GetByUserID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// UpdateUserLogs action: PUT /logs/:id
func (sc WorkLogController) UpdateUserLogs(c *gin.Context) {
	id := c.Params.ByName("id")
	var s model.WorkLog
	p, err := s.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
