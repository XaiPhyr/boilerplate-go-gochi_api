package tests

import (
	utils "gochi_api/utils"
	"os"
	"testing"
)

func TestMailer(t *testing.T) {
	os.Setenv("APP_ENVIRONMENT", "test")

	data := map[string]interface{}{
		"User": "RDev",
	}

	file := "../template/emails/welcome.html"
	if content, err := utils.ParseHTML(file, data); err == nil {
		utils.Mailer("Welcome!", content)
	}
}
