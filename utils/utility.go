package utils

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
	}

	ServerConfig struct {
		Endpoint string `yaml:"endpoint"`
		JwtKey   string `yaml:"jwt_key"`
	}

	DatabaseConfig struct {
		Host    string `yaml:"host"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Port    string `yaml:"port"`
		DB      string `yaml:"db"`
		SSLMode string `yaml:"sslmode"`
	}
)

func InitConfig() Config {
	var cfg Config

	f, err := os.Open("./conf/config.yml")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	return cfg
}

func ParseHTML(path string, data interface{}) (bodyHtml string, err error) {
	t := template.New(filepath.Base(path)).Funcs(template.FuncMap{})
	t, err = t.ParseFiles(path)

	if err != nil {
		fmt.Println("Error loading template", err.Error())
		return "", err
	} else {
		var tpl bytes.Buffer
		if err = t.Execute(&tpl, data); err == nil {
			bodyHtml = tpl.String()
		}
	}

	return
}

func InitDBConnect() *bun.DB {
	cfg := InitConfig()

	username := cfg.Database.User
	password := cfg.Database.Pass
	host := cfg.Database.Host
	port := cfg.Database.Port
	database := cfg.Database.DB
	sslmode := cfg.Database.SSLMode

	dsn := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=" + sslmode
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}
