package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type WorkLogController struct{}

/*************************************************
 *	specification;
 *	name 			= CreateUserLog
 *	Function 	= Create a User log
 *	note			= POST /logs
 *						= to create log need to jwt token
 *						= can create other user log
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.WorkLog JSON
 *  end of specification;
**************************************************/
func (sc WorkLogController) CreateUserLog(c *gin.Context) {
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

	var s model.WorkLog
	p, err := s.CreateLog(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= ShowUserLog
 *	Function 	= Get a self log
 *	note			= GET /logs
 *						= to get log need to jwt token
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.WorkLog JSON
 *  end of specification;
**************************************************/
func (sc WorkLogController) ShowUserLogs(c *gin.Context) {
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

	var s model.WorkLog
	p, err := s.GetByUserID(user_id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= UpdateUserLogs
 *	Function 	= Update a self log by id
 *	note			= PUT /logs/:id
 *						= to get log need to jwt token
 *						= can change other user log
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.WorkLog JSON
 *  end of specification;
**************************************************/
func (sc WorkLogController) UpdateUserLogs(c *gin.Context) {
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
	var s model.WorkLog
	p, err := s.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
