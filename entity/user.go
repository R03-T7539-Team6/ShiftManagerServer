/*
ここではモデルの定義だけ行い、
動作に関してはserviceディレクトリで管理する。
*/

package entity

// Uset is user models property
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
