package middlewares

import (
	"net/http"
	"strings"
	"time"

	s "gochi_api/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
)

type Middleware struct{}

var (
	LimiterExpiration = 1 * time.Minute
	RequestID         = middleware.RequestID
	RealIP            = middleware.RealIP
	Logger            = middleware.Logger
	Recoverer         = middleware.Recoverer
	GetHead           = middleware.GetHead
	HttpRate          = httprate.LimitByIP(100, LimiterExpiration)
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		t := strings.Split(auth, " ")[1]
		s.VerifyJWT(t, w)

		h.ServeHTTP(w, r)
	})
}

func UseMiddlewares(r *chi.Mux) {
	r.Use(RequestID)
	r.Use(RealIP)
	r.Use(Logger)
	r.Use(Recoverer)
	r.Use(GetHead)
	r.Use(HttpRate)
}
