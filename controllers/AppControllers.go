package controllers

import (
	"encoding/json"
	"gochi_api/models"
	"net/http"
)

type AppController struct {
	userModel *models.Users
}

func (a AppController) toJson(w http.ResponseWriter, b interface{}) {
	jsonMarshal, _ := json.MarshalIndent(b, "", "  ")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonMarshal))
}

func (a AppController) handleError(w http.ResponseWriter, errObj models.ErrorObject) {
	jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")

	w.WriteHeader(errObj.Code)
	w.Write([]byte(jsonMarshal))
}
