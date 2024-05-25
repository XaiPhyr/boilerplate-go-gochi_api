package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"

	u "gochi_api/utils"
)

type (
	MuxServer struct {
		Mux      *chi.Mux
		Endpoint string
	}

	JwtClaim struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	ErrorObject struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func Create(data interface{}) (result sql.Result, err error) {
	ctx := context.Background()

	db := u.InitDBConnect()
	q := db.NewInsert().Model(data)

	result, err = q.Exec(ctx)

	if err != nil {
		log.Printf("App Model \"Create\" %s", err)
	}

	return result, err
}
