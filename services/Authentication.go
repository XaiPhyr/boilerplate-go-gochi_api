package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gochi_api/models"
	u "gochi_api/utils"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	JwtExpiration        = 5 * time.Minute
	LimiterExpiration    = 1 * time.Minute
	JwtRefreshExpiration = 1 * time.Hour
)

var (
	JwtMethodHS256 = jwt.SigningMethodHS256
)

func GenerateJWT() (string, error) {
	app := u.InitConfig()
	jwtKey := app.Server.JwtKey

	claims := registerToken(JwtExpiration)

	token := jwt.NewWithClaims(JwtMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		log.Printf("Error: %s", err)
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(token string, w http.ResponseWriter) {
	app := u.InitConfig()
	jwtKey := app.Server.JwtKey

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	}

	t, err := jwt.ParseWithClaims(token, &models.JwtClaim{}, keyFunc)

	claims, _ := t.Claims.(*models.JwtClaim)

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Printf("Error: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
}

func RefreshJWT() {
	//
}

func HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(string(bytes))
	fmt.Println()
	fmt.Println()

	return nil
}

func CheckPassword() error {
	err := bcrypt.CompareHashAndPassword([]byte(""), []byte(""))

	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	return nil
}

func registerToken(duration time.Duration) (claims *models.JwtClaim) {
	claims = &models.JwtClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	return
}
