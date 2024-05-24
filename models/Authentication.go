package models

import "time"

type (
	Authentication struct {
		Username     string `json:"username"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		RememberMe   bool   `json:"remember_me"`
	}

	Login struct {
		Username   string `gorm:"column:username" json:"username"`
		Password   string `gorm:"column:password" json:"password"`
		UserType   string `gorm:"column:user_type" json:"user_type"`
		RememberMe bool   `gorm:"-" json:"remember_me"`
	}

	Register struct {
		*Login
		Email      string    `gorm:"column:email" json:"email"`
		Status     string    `gorm:"column:status" json:"status"`
		Active     bool      `gorm:"column:active" json:"active"`
		CreatedAt  time.Time `gorm:"column:created_at"`
		ModifiedAt time.Time `gorm:"column:modified_at"`
	}
)

func (Register) TableName() string {
	return "users"
}

func NewRegister() *Register {
	return &Register{}
}
