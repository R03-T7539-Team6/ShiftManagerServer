package model

// for Login, JWT authentication
type Authorization struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

// for Response jwt token
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user" gorm:"foreignKey:UserID"`
}

// for get Authorization header token
type AuthorizationHeader struct {
	Authorization string `header:"Authorization"`
}
