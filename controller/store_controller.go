package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/gin-gonic/gin"
)

type StoreController struct{}

/*************************************************
 *	specification;
 *	name 			= CreateStore
 *	Function 	= Create a Store file
 *	note			= Store is Unique
 *						= POST /stores
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.20
 *  input 		= gin.Context
 *  output    = JSON with Status code
 * 						= model.Store JSON
 *  end of specification;
**************************************************/
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

	var s model.Store
	p, err := s.CreateStore(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= ShowStore
 *	Function 	= Get a Store file
 *	note			= id is store_id, and remove user password
 *						= POST /stores/:id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.20
 *  input 		= gin.Context
 *  output    = JSON with Status code
 * 						= model.Store JSON
 *  end of specification;
**************************************************/
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
