package config

import (
	"authentication_backend/internal"
	"authentication_backend/var/email"
	"os"

	"github.com/wneessen/go-mail"
)

func InitEmail() {

	email.Contact = internal.NewEmailClient(
		"contact@example.com",
		os.Getenv("EMAIL_HOST"),
		os.Getenv("EMAIL_PORT"),
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
		mail.SMTPAuthPlain,
		mail.NoTLS)
}
