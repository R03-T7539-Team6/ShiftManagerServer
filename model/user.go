package model

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/R03-T7539-Team6/ShiftManagerSerer/utility"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Status value
/*
	Normal          	// 通常
	InLeaveOfAbsence	// 休職中
	Retired         	// 退職済み
	NotHired        	// 未採用
	Others          	// その他
*/
type status string

const (
	Normal           status = "Normal"
	InLeaveOfAbsence status = "InLeaveOfAbsence"
	Retired          status = "Retired"
	NotHired         status = "NotHired"
	Others           status = "Others"
)

// Group value
/*
	None,										// デフォルト値
	SystemAdmin							// システム管理者（データ編集用)
	SuperUser							　// 権限ユーザ（店舗責任者）
	NormalUser							// 一般ユーザ
	ForTimeRecordTerminal		// 勤怠登録/店舗端末用
*/

type group string

const (
	None                  group = "None"
	SystemAdmin           group = "SystemAdmin"
	SuperUser             group = "SuperUser"
	NormalUser            group = "NormalUser"
	ForTimeRecordTerminal group = "ForTimeRecordTerminal"
)

/*
	// gorm.Modelの定義
	type Model struct {
  	ID        uint           `gorm:"primaryKey"`
  	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt gorm.DeletedAt `gorm:"index"`
	}

*/

// Uset is user models property
type User struct {
	gorm.Model
	UserID    string `json:"user_id" gorm:"unique"`
	Password  string `json:"password"`
	StoreID   string `json:"store_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserState status `json:"user_state"`
	UserGroup group  `json:"user_group"`
}

/*************************************************
 *	specification;
 *	name 			= CreateModel
 *	Function 	= create User model in user table
 *	note			= user table is related json
 *						= make password hashed
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context http.request
 *  output    = User: User sturct
 * 						= error value
 *  end of specification;
**************************************************/
func (s User) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	// パスワードはハッシュ化して保存する
	u.Password = utility.HashStr(u.Password, "sha256")

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByID
 *	Function 	= get a User model in user table
 *	note			= user table is related json
 *						= searching by user_id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= id: string, this is user_id
 *  output    = User: User sturct
 * 						= error value
 *  end of specification;
**************************************************/
func (s User) GetByID(id string) (User, error) {
	db := db.GetDB()
	var u User
	if err := db.Where("user_id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByIDWithPassword
 *	Function 	= get a User model in user table with password
 *	note			= user table is related json
 *						= searching by user_id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= id: string, this is user_id
 *  output    = User: User sturct
 * 						= error value
 *  end of specification;
**************************************************/
func (s User) GetByIDWithPassword(id string) (User, error) {
	db := db.GetDB()
	var u User
	if err := db.Where("user_id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

/*************************************************
 *	specification;
 *	name 			= UpdateByID
 *	Function 	= Update a User model in user table
 *	note			= user table is related json
 *						= searching by user_id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= id: string, this is user_id
 * 						= c: *gin.Context http.request
 *  output    = User: User sturct
 * 						= error value
 *  end of specification;
**************************************************/
func (s User) UpdateByID(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("user_id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	// パスワードの更新なのかを確かめる
	var has_password map[string]interface{}
	c.BindJSON(&has_password)
	if _, ok := has_password["password"]; ok {
		u.Password = utility.HashStr(u.Password, "sha256")
	}

	db.Save(&u)

	return u, nil
}

/*************************************************
 *	specification;
 *	name 			= DeleteByID
 *	Function 	= delete a User model in user table
 *	note			= user table is related json
 *						= searching by user_id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= id: string, this is user_id
 *  output    = error value
 *  end of specification;
**************************************************/
func (s User) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("user_id = ?", id).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
