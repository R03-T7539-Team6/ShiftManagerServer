package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/R03-T7539-Team6/ShiftManagerSerer/utility"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AuthorizationController struct{}

func (ac AuthorizationController) Login(c *gin.Context) {
	var a model.Authorization
	var u model.User

	// Body取り出し、JSONにバインドする。
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	// ユーザー認証する
	user_id_auth := a.UserID
	user_password_auth := utility.HashStr(a.Password, "sha256")

	user_in_database, err := u.GetByID(user_id_auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Wrong User ID")
		fmt.Println(err)
	}

	// compare the from the request, with the one we defined:
	if user_in_database.UserID != user_id_auth || user_in_database.Password != user_password_auth {
		c.JSON(http.StatusUnauthorized, "Wrong Password")
		return
	}
	log.Println("login success: user id = ", user_in_database.UserID)

	token, err := CreateToken(user_in_database.UserID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, &model.LoginResponse{
		Token: token,
	})
}

func CreateToken(user_id string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	secretkey := os.Getenv("ACCESS_SECRET")
	// create token
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// setting claims
	token.Claims = jwt.MapClaims{
		"authorized": true,
		"user":       user_id,
		"exp":        time.Now().Add(time.Hour * 3).Unix(), // 有効期限を設定: 3時間
	}

	// 署名
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verifytoken(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	secretkey := os.Getenv("ACCESS_SECRET")
	//jwtの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
