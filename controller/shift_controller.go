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
	var sr model.Shift
	p, err := sr.CreateShift(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// ShowShift action: GET /shift
// 									 GET /shift/?is_reuqest=true
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

	var p []model.Shift
	if is_request != "" {
		is_request_parse, _ := strconv.ParseBool(is_request)
		p, err = sr.GetByUserIdAndIsRequest(user_id, is_request_parse)
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
	// ペイロード読み出し
	// claims := token.Claims.(jwt.MapClaims)
	// user_id := fmt.Sprintf("%s", claims["user"])

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

// ShowShift action: GET /shift
// 									 GET /shift?is_reuqest=true&work_date=2021-01-02
// func (sc ShiftController) ShowShift(c *gin.Context) {
// 	// headerを取得
// 	h := model.AuthorizationHeader{}
// 	if err := c.ShouldBindHeader(&h); err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
// 		return
// 	}
// 	tokenString := h.Authorization
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	// tokenの認証
// 	_, err := Verifytoken(tokenString)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
// 		return
// 	}
// 	// ペイロード読み出し
// 	// claims := token.Claims.(jwt.MapClaims)
// 	// user_id := fmt.Sprintf("%s", claims["user"])
// 	var sr model.Shift
// 	work_date := c.Query("work_date")
// 	is_request := c.Query("is_request")

// 	var p []model.Shift

// 	// layout := "2000-01-01T01:01:01Z"
// 	work_date_parse, _ := time.Parse(time.RFC3339Nano, work_date)
// 	is_request_parse, _ := strconv.ParseBool(is_request)
// 	switch {
// 	case work_date != "" && is_request != "":
// 		p, err = sr.GetByWorkDateAndIsRequest(work_date_parse, is_request_parse)
// 	case work_date != "":
// 		p, err = sr.GetByWorkDate(work_date_parse)
// 	case is_request != "":
// 		p, err = sr.GetByIsRequest(is_request_parse)
// 	default:
// 		p, err = nil, errors.New("record not found")
// 	}

// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, p)
// 	}
// }

// DeleteShift action: DELETE /shift/:id
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
	// ペイロード読み出し
	// claims := token.Claims.(jwt.MapClaims)
	// user_id := fmt.Sprintf("%s", claims["user"])
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

// DeleteRequest action: DELETE /shift/requests/:id
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
	// ペイロード読み出し
	// claims := token.Claims.(jwt.MapClaims)
	// user_id := fmt.Sprintf("%s", claims["user"])
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

// ****************ShiftSchedule
// CreateSchedule action: POST /shift/schedule
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
	// claims := token.Claims.(jwt.MapClaims)
	// user_id := fmt.Sprintf("%s", claims["user"])
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
	var ss model.ShiftSchedule
	p, err := ss.GetByStoreId(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
