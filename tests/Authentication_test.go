package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	c "gochi_api/controllers"
	"gochi_api/models"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

var (
	APIURL = "/api/login"
	Token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY1Mzg5OTl9.nUjkEfIYzdHX1xV-bR4GrneaMvWqyjdmqW64hBn-De8"
)

func InitAuthenticationTest(req *http.Request) *httptest.ResponseRecorder {
	os.Setenv("APP_ENVIRONMENT", "test")

	a := &c.Authentication{}
	r := chi.NewRouter()

	var mux = models.MuxServer{
		Mux:      r,
		Endpoint: "/api",
	}

	a.InitAuthentication(mux)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	return rr
}

func TestAuthenticationLogin(t *testing.T) {
	jsonBody := map[string]interface{}{
		"username": "rdev",
		"password": "iamsuperadmin",
	}

	b, _ := json.Marshal(jsonBody)

	req, _ := http.NewRequest("GET", APIURL, nil)
	req.Header.Set("Authentication", string(b))
	rr := InitAuthenticationTest(req)

	require.Equal(t, http.StatusOK, rr.Code)
}

func TestAuthenticationRegister(t *testing.T) {
	jsonBody := map[string]interface{}{
		"email":     "rdev@local",
		"username":  "rdev",
		"password":  "iamsuperadmin",
		"user_type": "superadmin",
	}

	b, _ := json.Marshal(jsonBody)

	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(string(b)))
	rr := InitAuthenticationTest(req)

	require.Equal(t, http.StatusOK, rr.Code)
}

func TestPageNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/404", nil)
	rr := InitAuthenticationTest(req)

	require.Equal(t, 404, rr.Code)
}
