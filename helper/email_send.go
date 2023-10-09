package helper

import (
	"fmt"
	"go-schedule-email-task/database"
	"go-schedule-email-task/model"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func sendEmail(smtpServer, smtpPort, senderEmail, senderPassword string, msg []byte, to []string) error {
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, to, msg)
	return err
}

func composeEmail(from string, to []string, subject, body string) string {
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = strings.Join(to, ",")
	headers["Subject"] = subject

	var email strings.Builder
	for key, value := range headers {
		email.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	email.WriteString("\r\n" + body)
	return email.String()
}

func sendingEmail(receivingPerson string) error {
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{receivingPerson}

	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	subject := "AUTOMATION EMAIL"
	body := "This is the reminder Email."

	msg := []byte(composeEmail(from, to, subject, body))
	err := sendEmail(smtpHost, smtpPort, from, password, msg, to)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}

func SendingEmailOperation() {

	var emp []model.Employee

	database.Database.Find(&emp)
	for _, each := range emp {
		if each.EmailSentTime.Before(time.Now()) && each.EmailSend == false {
			fmt.Printf("Sending Email to %s\n", each.Email)
			err := sendingEmail(each.Email)
			if err != nil {
				fmt.Println("Error Occurred ->", err)
				return
			}
			each.EmailSend = true
			database.Database.Save(&each)

		}
	}
}
