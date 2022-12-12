package utils

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/gofrs/uuid"
)

func SendMail(token uuid.UUID, email string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	toList := []string{email}
	host := "smtp.gmail.com"
	port := "587"

	fromm := fmt.Sprintf("From: <%s>\r\n", from)
	to := fmt.Sprintf("To: <%s>\r\n", email)
	subject := "Subject: Survivor Coders Change Password\r\n"
	content := "This is the Link for password: \r\n http://localhost:8000/createpassword/" + token.String() + "\r\n\n\nSurvivor Coders\r\n"

	msg := fromm + to + subject + "\r\n" + content

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	return err
}
