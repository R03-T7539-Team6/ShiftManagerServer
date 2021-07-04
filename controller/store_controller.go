package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/gin-gonic/gin"
)

type StoreController struct{}

// Create action: POST /stores
func (sc StoreController) CreateStore(c *gin.Context) {
	// headerを取得
	h := model.AuthorizationHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	tokenString := h.Authorization
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// tokenの認証
	_, err := Verifytoken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	// ペイロード読み出し
	// claims := token.Claims.(jwt.MapClaims)
	// user_id := fmt.Sprintf("%s", claims["user"])

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
	// headerを取得
	h := model.AuthorizationHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	tokenString := h.Authorization
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// tokenの認証
	_, err := Verifytoken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	// ペイロード読み出し
	// claims := token.Claims.(jwt.MapClaims)
	// user_id := fmt.Sprintf("%s", claims["user"])
	id := c.Params.ByName("id")
	var s model.Store
	p, err := s.GetByStoreID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		for i := 0; i < len(p.Worker); i++ {
			p.Worker[i].Password = ""
		}
		c.JSON(200, p)
	}
}
