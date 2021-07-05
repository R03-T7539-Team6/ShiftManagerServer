package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

/*************************************************
 *	specification;
 *	name 			= Create
 *	Function 	= Create a User
 *	note			= POST /Users
 *						= to create user need to jwt token
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.11/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.User JSON without password
 *  end of specification;
**************************************************/
func (pc UserController) Create(c *gin.Context) {
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

	var u model.User
	p, err := u.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		p.Password = ""
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= Show
 *	Function 	= Get self User info
 *	note			= GET /Users
 *						= to get user need to jwt token
 *						= can not get other user info
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.11/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.User JSON without password
 *  end of specification;
**************************************************/
func (pc UserController) Show(c *gin.Context) {
	// headerを取得
	h := model.AuthorizationHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	tokenString := h.Authorization
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// tokenの認証
	token, err := Verifytoken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	// ペイロード読み出し
	claims := token.Claims.(jwt.MapClaims)
	user_id := fmt.Sprintf("%s", claims["user"])

	// id := c.Params.ByName("id")
	var u model.User
	p, err := u.GetByID(user_id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		p.Password = ""
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= Update
 *	Function 	= Update self User info
 *	note			= PUT /Users
 *						= to put user need to jwt token
 *						= can not change other user
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.11/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.User JSON without password
 *  end of specification;
**************************************************/
func (pc UserController) Update(c *gin.Context) {
	// headerを取得
	h := model.AuthorizationHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	tokenString := h.Authorization
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// tokenの認証
	token, err := Verifytoken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	// ペイロード読み出し
	claims := token.Claims.(jwt.MapClaims)
	user_id := fmt.Sprintf("%s", claims["user"])
	// id := c.Params.ByName("id")
	var u model.User
	p, err := u.UpdateByID(user_id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		p.Password = ""
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= Delete
 *	Function 	= Delete self User info
 *	note			= DELETE /Users
 *						= to delete user need to jwt token
 *						= when delete self, expire token
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.11/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= deleted userid JSON
 *  end of specification;
**************************************************/
func (pc UserController) Delete(c *gin.Context) {
	// headerを取得
	h := model.AuthorizationHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	tokenString := h.Authorization
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// tokenの認証
	token, err := Verifytoken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	// ペイロード読み出し
	claims := token.Claims.(jwt.MapClaims)
	user_id := fmt.Sprintf("%s", claims["user"])
	// id := c.Params.ByName("id")
	var u model.User

	if err := u.DeleteByID(user_id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + user_id: "deleted",
		})
	}
}
