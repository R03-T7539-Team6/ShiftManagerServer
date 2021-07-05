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
)

type AuthorizationController struct{}

/*************************************************
 *	specification;
 *	name 			= Signup
 *	Function 	= Signup handler
 *	note			= Signup is add user without auth
							= POST /signup
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00
 *  input 		= gin.Context
 *  output    = JSON with Status code
 * 						= model.User JSON without password
 *  end of specification;
**************************************************/
func (pc AuthorizationController) Signup(c *gin.Context) {
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
 *	name 			= Login
 *	Function 	= Login handler
 *	note			= Login is get token for JWT Authenticate
							= compare user_name and password in DB
							= POST /login
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10/V1.11
 *  input 		= gin.Context
 *  output    = JSON with Status code
 * 						= model.User JSON without password
 *  end of specification;
**************************************************/
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

	user_in_database, err := u.GetByIDWithPassword(user_id_auth)
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

/*************************************************
 *	specification;
 *	name 			= CreateToken
 *	Function 	= Create token for JWT
 *	note			= When the token create, limite the token 3 hours.
							= need to save token in frontend
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00
 *  input 		= user_id: string
 *  output    = jwt_token: string
 *  end of specification;
**************************************************/
func CreateToken(user_id string) (string, error) {
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

/*************************************************
 *	specification;
 *	name 			= Verifytoken
 *	Function 	= Verify jwt token
 *	note			= Use SECRET_KEY in .ENV
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00
 *  input 		= jwt_token : string
 *  output    = Parsing token string: string
 *  end of specification;
**************************************************/
func Verifytoken(tokenString string) (*jwt.Token, error) {
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
