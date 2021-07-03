/*
ここではモデルの定義だけ行い、
動作に関してはserviceディレクトリで管理する。
*/
package model

type Authorization struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AuthorizationHeader struct {
	Authorization string `header:"Authorization"`
}
