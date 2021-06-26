/*
ここではモデルの定義だけ行い、
動作に関してはserviceディレクトリで管理する。
*/

package entity

import "github.com/jinzhu/gorm"

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
	UserID    uint      `json:"user_id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserState UserState `json:"user_state"`
	UserGroup UserGroup `json:"user_group"`
}
