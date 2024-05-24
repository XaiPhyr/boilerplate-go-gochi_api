package controllers

import (
	"encoding/json"
	"fmt"
	"gochi_api/models"
	"net/http"

	mw "gochi_api/middlewares"

	"github.com/go-chi/chi"
)

type HelloWorld struct {
	AppController
}

func (hw HelloWorld) InitHelloWorld(m models.MuxServer) {
	m.Mux.Route(m.Endpoint+"/hello", func(r chi.Router) {
		r.Get("/", hw.GetAll)

		r.Group(func(r chi.Router) {
			r.Use(mw.Authenticate)
			r.Post("/", hw.Post)
			r.Delete("/", hw.DeleteAll)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", hw.GetOne)
				r.Patch("/", hw.Patch)
				r.Put("/", hw.Put)
				r.Delete("/", hw.DeleteOne)
			})
		})
	})
}

func (hw HelloWorld) GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET ALL"))
}

func (hw HelloWorld) GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("GET ONE ID: %v", id)))
}

func (hw HelloWorld) Post(w http.ResponseWriter, r *http.Request) {
	var data models.HelloWorld
	err := json.NewDecoder(r.Body).Decode(&data)
	token := r.Header.Get("Authorization")

	if err != nil {
		return
	}

	jsonMarshal, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println()
	fmt.Println("POST AUTHORIZATION: " + token)
	fmt.Println("DECODER: ", fmt.Sprintf("%v", string(jsonMarshal)))
	fmt.Println()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("POST"))
}

func (hw HelloWorld) Patch(w http.ResponseWriter, r *http.Request) {
	var data models.HelloWorld
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return
	}

	jsonMarshal, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println()
	fmt.Println("PATCH : " + id)
	fmt.Println("DECODER: ", fmt.Sprintf("%v", string(jsonMarshal)))
	fmt.Println()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("PATCH ONE ID: %v", id)))
}

func (hw HelloWorld) Put(w http.ResponseWriter, r *http.Request) {
	var data models.HelloWorld
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return
	}

	jsonMarshal, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println()
	fmt.Println("PUT : " + id)
	fmt.Println("DECODER: ", fmt.Sprintf("%v", string(jsonMarshal)))
	fmt.Println()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("UPDATE ONE ID: %v", id)))
}

func (hw HelloWorld) DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("DELETE ALL"))
}

func (hw HelloWorld) DeleteOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("DELETE ONE ID: %v", id)))
}
