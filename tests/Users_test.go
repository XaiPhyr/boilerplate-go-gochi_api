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
	req.Header.Set("Authorization", "Bearer 1")
	rr := InitUsersTest(req)

	require.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateUser(t *testing.T) {
	uuid := "09a0211f-1284-458b-82d7-2c2d7ece39f0"
	url := "/api/users/" + uuid + "/update"

	jsonBody := map[string]interface{}{
		"id":       52,
		"username": "rdev6update1",
		"email":    "upu@example.com",
	}

	b, _ := json.Marshal(jsonBody)

	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(b)))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsImV4cCI6MTcxNjc5MTM0NX0.czTEb3XKPKOlVXLftigu1wlpt0rYeM6Zb8VnPZKVSrY")
	rr := InitUsersTest(req)

	require.Equal(t, http.StatusOK, rr.Code)
}
