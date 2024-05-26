package controllers

import (
	"encoding/json"
	"gochi_api/middlewares"
	"gochi_api/models"
	"net/http"
)

type AppController struct {
	mw        middlewares.Middleware
	userModel *models.Users
}

func (a AppController) toJson(w http.ResponseWriter, b interface{}) {
	jsonMarshal, _ := json.MarshalIndent(b, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonMarshal))
}

func (a AppController) handleError(w http.ResponseWriter, code int, message string) {
	errObj := models.ErrorObject{
		Code:    code,
		Message: message,
	}

	jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errObj.Code)
	w.Write([]byte(jsonMarshal))
}
