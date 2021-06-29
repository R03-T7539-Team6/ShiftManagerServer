package controller

import (
	"fmt"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/gin-gonic/gin"
)

// UserController is User UserController
type UserController struct{}

// Index action: GET /Users
func (pc UserController) Index(c *gin.Context) {
	var u model.User
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /Users
func (pc UserController) Create(c *gin.Context) {
	var u model.User
	p, err := u.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Show action: GET /Users/:id
func (pc UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var u model.User
	p, err := u.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /Users/:id
func (pc UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var u model.User
	p, err := u.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /Users/:id
func (pc UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var u model.User

	if err := u.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + id: "deleted",
		})
	}
}
