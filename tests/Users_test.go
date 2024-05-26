package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	c "gochi_api/controllers"
	"gochi_api/models"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

func InitUsersTest(req *http.Request) *httptest.ResponseRecorder {
	os.Setenv("APP_ENVIRONMENT", "test")

	a := &c.Users{}
	r := chi.NewRouter()

	var mux = models.MuxServer{
		Mux:      r,
		Endpoint: "/api",
	}

	a.InitUsers(mux)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	return rr
}

func TestGetAllUsers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/users", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsImV4cCI6MTcxNjc0MTU4MX0.4jt-iL4V2ncKky9kayDXRQ8Ac5hB8htAnjbLv8MmX2Y")
	rr := InitUsersTest(req)

	require.Equal(t, http.StatusOK, rr.Code)
}
