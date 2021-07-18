package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ShiftController struct{}

/*************************************************
 *	specification;
 *	name 			= CreateShift
 *	Function 	= Create a shift
 *	note			= POST /shifts
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.11/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.Shift JSON
 *  end of specification;
**************************************************/
func (sc ShiftController) CreateShift(c *gin.Context) {
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

	var sr model.Shift
	p, err := sr.CreateShift(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= ShowShiftByUser
 *	Function 	= get a user by user id or query
 *	note			= GET /shifts
							= GET /shifts?is_request=true
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.20/V1.30
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.Shift JSON
 *  end of specification;
**************************************************/
func (sc ShiftController) ShowShiftByUser(c *gin.Context) {
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
	var sr model.Shift
	// id := c.Params.ByName("id")
	is_request := c.Query("is_request")
	work_date := c.Query("target_date")

	var p []model.Shift
	if work_date == "" && is_request != "" {
		is_request_parse, _ := strconv.ParseBool(is_request)
		p, err = sr.GetByUserIdAndIsRequest(user_id, is_request_parse)
	} else if work_date != "" && is_request == "" {
		// work_date_parse, _ := time.Parse("yyyy-mm-ddT00-00-00Z", work_date)
		p, err = sr.GetByUserIDandWorkDate(user_id, work_date)
	} else if work_date != "" && is_request != "" {
		is_request_parse, _ := strconv.ParseBool(is_request)
		p, err = sr.GetByWorkDateAndIsRequest(work_date, is_request_parse)
	} else {
		p, err = sr.GetByUserId(user_id)
	}

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= UpdateShift
 *	Function 	= update a shift by id
 *	note			= PUT /shifts/:id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.Shift JSON
 *  end of specification;
**************************************************/
// Update action: POST /shift/:id
func (sc ShiftController) UpdateShift(c *gin.Context) {
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
	var sr model.Shift
	p, err := sr.UpdateShift(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= DeleteShift
 *	Function 	= Delete a shift by id
 *	note			= DELETE /shifts/:id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= deleted id
 *  end of specification;
**************************************************/
func (sc ShiftController) DeleteShift(c *gin.Context) {
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
	var sr model.Shift
	if err := sr.DeleteById(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + id: "deleted",
		})
	}
}

/*************************************************
 *	specification;
 *	name 			= CreateRequest
 *	Function 	= Create user request file
 *	note			= POST /shifts/requests
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.ShiftRequest JSON
 *  end of specification;
**************************************************/
func (sc ShiftController) CreateRequest(c *gin.Context) {
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
	var sr model.ShiftRequest
	p, err := sr.CreateShiftRequest(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= ShowRequest
 *	Function 	= Get a shift request file by user_id
 *	note			= GET /shifts/requests
 *						= user_id gets a JWT token
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.ShiftRequest JSON
 *  end of specification;
**************************************************/
func (sc ShiftController) ShowRequest(c *gin.Context) {
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
	var sr model.ShiftRequest
	p, err := sr.GetByUserId(user_id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= DeleteRequest
 *	Function 	= Delete a request by id
 *	note			= DELETE /shifts/requests/:id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= deleted id JSON
 *  end of specification;
**************************************************/
func (sc ShiftController) DeleteRequest(c *gin.Context) {
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
	var sr model.ShiftRequest
	if err := sr.DeleteById(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{
			"id #" + id: "deleted",
		})
	}
}

/*************************************************
 *	specification;
 *	name 			= CreateSchedule
 *	Function 	= Create a schedule
 *	note			= POST /shifts/schedule
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.ShiftSchedule JSON
 *  end of specification;
**************************************************/
func (sc ShiftController) CreateSchedule(c *gin.Context) {
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
	var ss model.ShiftSchedule
	p, err := ss.CreateShiftSchedule(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

/*************************************************
 *	specification;
 *	name 			= ShowSchedule
 *	Function 	= Get a schedule by store_id
 *	note			= GET /shifts/schedule/:id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.20
 *  input 		= gin.Context
 *  output    = JSON with status code
 * 						= model.ShiftSchedule JSON
 *  end of specification;
**************************************************/
func (sc ShiftController) ShowSchedule(c *gin.Context) {
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
	id := c.Params.ByName("id")
	work_date := c.Query("target_date")

	var ss model.ShiftSchedule
	var p model.ShiftSchedule
	if work_date != "" {
		p, err = ss.GetByStoreIdAndTargetDate(id, work_date)
	} else {
		p, err = ss.GetByStoreId(id)
	}

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
