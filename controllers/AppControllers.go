package controllers

import (
	"encoding/json"
	"net/http"
)

type AppController struct{}

func (a AppController) toJson(w http.ResponseWriter, b interface{}) {
	jsonMarshal, _ := json.MarshalIndent(b, "", "  ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonMarshal))
}
