package controller

import (
	"fmt"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/gin-gonic/gin"
)

type StoreController struct{}

// Create action: POST /stores
func (sc StoreController) CreateStore(c *gin.Context) {
	var s model.Store
	p, err := s.CreateStore(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Show action: POST /stores/:id
func (sc StoreController) ShowStore(c *gin.Context) {
	id := c.Params.ByName("id")
	var s model.Store
	p, err := s.GetByStoreID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
