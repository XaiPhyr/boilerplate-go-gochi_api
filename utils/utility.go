package utils

import (
	"bytes"
	"fmt"
	"gochi_api/models"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func InitConfig() models.Config {
	var cfg models.Config

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
