/*
ここではモデルの定義だけ行い、
動作に関してはserviceディレクトリで管理する。
*/
package model

import "github.com/jinzhu/gorm"

type Authorization struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Password string `json:"password"`
}
