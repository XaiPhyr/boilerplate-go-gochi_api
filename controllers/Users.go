package controllers

import (
	"gochi_api/models"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Users struct {
	AppController
}

func (u Users) InitUsers(m models.MuxServer) {
	m.Mux.Route(m.Endpoint+"/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(u.mw.Authenticate)

			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/detail", u.GetUser)
			})
		})
	})
}

func (u Users) GetUser(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	result, err := u.userModel.ReadOneUser(uuid)

	if err != nil {
		log.Printf("Error: %s", err)
		u.handleError(w, http.StatusNotFound, "User not found")
		return
	}

	u.toJson(w, result)
}