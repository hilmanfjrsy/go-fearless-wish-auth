package datatransfers

import (
	"time"
)

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Email string `json:"email" binding:"-"`
	Phone string `json:"phone" binding:"-"`
}

type UserInfo struct {
	Username  string    `uri:"username" json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
