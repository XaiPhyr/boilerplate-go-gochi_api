package models

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (
	UserResults struct {
		Count int      `json:"count"`
		Data  *[]Users `json:"data"`
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

func (u *Users) ReadAllUsers(limit, offset int) (data *[]Users, count int, err error) {
	data = new([]Users)

	q, ctx := Read(data)
	q = q.ExcludeColumn("password")
	q = q.Limit(limit)
	q = q.Offset(offset)
	q = q.Order("created_at DESC")

	count, err = q.ScanAndCount(ctx)
	return
}

func (u *Users) ReadOneUser(uuid string) (data *Users, err error) {
	data = new(Users)

	q, ctx := Read(data)
	q = q.ExcludeColumn("password")
	q = q.Where("uuid = ?", uuid)

	err = q.Scan(ctx)
	return
}

func (u *Users) CreateUser(w http.ResponseWriter, data *Users, fn func(w http.ResponseWriter, code int, message string)) (sql.Result, error) {
	q, ctx := Create(data)

	result, err := q.Exec(ctx)

	if err != nil {
		u.HandleUserError(w, err, fn)
		return nil, err
	}

	return result, nil
}

func (u *Users) UpdateOneUser(data *Users) (err error) {
	return
}

func (u *Users) DeleteAllUsers(data *Users) (err error) {
	return
}

func (u *Users) DeleteOneUser(data *Users) (err error) {
	return
}

func (u *Users) ReadByUsername(user string) (data *Users, err error) {
	data = new(Users)

	q, ctx := Read(data)
	q = q.Where("username = ?", user)

	err = q.Scan(ctx)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	return
}

func (u *Users) HandleUserError(w http.ResponseWriter, err error, fn func(w http.ResponseWriter, code int, message string)) {
	log.Printf("Error: %s", err)

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

			case "username_alphanumeric_check":
				errObj = ErrorObject{
					Code:    http.StatusBadRequest,
					Message: "Username must be aplhanumeric",
				}
			}
		}
	}

	fn(w, errObj.Code, errObj.Message)
}
