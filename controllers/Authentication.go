package controllers

import (
	"encoding/json"
	"gochi_api/models"
	"log"
	"net/http"
	"strings"

	s "gochi_api/services"

	"github.com/go-chi/chi"
)

type Authentication struct {
	AppController
}

func (a Authentication) InitAuthentication(m models.MuxServer) {
	m.Mux.Route(m.Endpoint+"/login", func(r chi.Router) {
		r.Get("/", a.Login)
	})

	m.Mux.Route(m.Endpoint+"/register", func(r chi.Router) {
		r.Post("/", a.Register)
	})
}

func (a Authentication) Login(w http.ResponseWriter, r *http.Request) {
	var info models.Authentication
	var l models.Login

	token, err := s.GenerateJWT()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auth := r.Header.Get("Authentication")
	readAuth := strings.NewReader(auth)
	err = json.NewDecoder(readAuth).Decode(&l)

	s.HashPassword(l.Password)

	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	info.Token = token
	info.Username = l.Username
	info.RememberMe = l.RememberMe

	a.toJson(w, info)
}

func (a Authentication) Register(w http.ResponseWriter, r *http.Request) {
	user := a.userModel.NewRegister()

	if err := json.NewDecoder(r.Body).Decode(&user); err == nil {
		hashPassword, _ := s.HashPassword(user.Password)
		user.Password = hashPassword

		err := a.userModel.CreateUser(user)

		if err != nil {
			errObj := a.userModel.HandleAuthenticationError(err)
			a.handleError(w, errObj)
			return
		}

		a.toJson(w, user)
	}
}
