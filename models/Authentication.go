package models

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (
	Authentication struct {
		Username     string `json:"username"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		RememberMe   bool   `json:"remember_me"`
	}

	Login struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		UserType   string `json:"user_type"`
		RememberMe bool   `json:"remember_me"`
	}

	Users struct {
		bun.BaseModel `bun:"table:users,alias:u"`

		ID        int64     `bun:"id,pk,autoincrement" json:"id"`
		Username  string    `bun:"username" json:"username"`
		Password  string    `bun:"password" json:"password"`
		UserType  string    `bun:"user_type" json:"user_type"`
		Email     string    `bun:"email" json:"email"`
		UUID      string    `bun:"uuid" json:"uuid"`
		Status    string    `bun:"status,default:O" json:"status"`
		Active    bool      `bun:"active" json:"active"`
		CreatedAt time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
		UpdatedAt time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
		DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero" json:"deleted_at"`
	}
)

func (u *Users) NewRegister() *Users {
	return &Users{
		UUID:   uuid.Must(uuid.NewRandom()).String(),
		Status: "O",
		Active: false,
	}
}

func (u *Users) CreateUser(data *Users) error {
	_, err := Create(data)
	return err
}

func (u *Users) HandleAuthenticationError(err error) ErrorObject {
	errObj := ErrorObject{
		Code:    400,
		Message: "Bad Request",
	}

	if err != nil {
		pgErr := err.(pgdriver.Error)

		if pgErr.IntegrityViolation() {
			log.Printf("pgErr integrity violation: %s", pgErr.Field('n'))

			switch pgErr.Field('n') {
			case "username_unique_idx":
				errObj = ErrorObject{
					Code:    http.StatusBadRequest,
					Message: "Username already taken",
				}
			}
		}
	}

	return errObj
}
