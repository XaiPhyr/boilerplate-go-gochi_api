package models

import (
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
)

type (
	Config struct {
		Server ServerConfig `yaml:"server"`
	}

	ServerConfig struct {
		Endpoint string `yaml:"endpoint"`
		JwtKey   string `yaml:"jwt_key"`
	}

	DatabaseConfig struct{}

	MuxServer struct {
		Mux      *chi.Mux
		Endpoint string
	}

	JwtClaim struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}
)
