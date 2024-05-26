package routers

import (
	"encoding/json"
	c "gochi_api/controllers"
	"gochi_api/models"
	"log"
	"net/http"

	mw "gochi_api/middlewares"
	utils "gochi_api/utils"

	"github.com/go-chi/chi"
)

var (
	auth = &c.Authentication{}
	user = &c.Users{}
)

func NewRoutes() chi.Router {
	api := utils.InitConfig()

	r := chi.NewRouter()
	mw.UseMiddlewares(r)

	var mux = models.MuxServer{
		Mux:      r,
		Endpoint: api.Server.Endpoint,
	}

	// @routes
	auth.InitAuthentication(mux)
	user.InitUsers(mux)

	// @status 404, 405
	PageNotFound(r)
	MethodNotAllowed(r)

	return r
}

func PageNotFound(r *chi.Mux) {
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		content, err := utils.ParseHTML("template/404.html", nil)

		if err != nil {
			log.Printf("Error: %s", err)
			return
		}

		w.Write([]byte(string(content)))
	})
}

func MethodNotAllowed(r *chi.Mux) {
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		b := map[string]interface{}{
			"err":  "method not allowed",
			"code": http.StatusMethodNotAllowed,
		}

		jsonMarshal, _ := json.MarshalIndent(b, "", "  ")

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonMarshal))
	})
}
