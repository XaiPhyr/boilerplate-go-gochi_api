package controllers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type AppController struct{}

func (a AppController) toJson(w http.ResponseWriter, b interface{}) {
	jsonMarshal, _ := json.MarshalIndent(b, "", "  ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonMarshal))
}

func (a AppController) InitDBConnect() (*gorm.DB, error) {
	return nil, nil
}
