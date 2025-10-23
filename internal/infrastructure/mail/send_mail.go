package mail

import (
	"emailn/internal/domain/campaign"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(campaing *campaign.Campaign) error {

	dialer := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("GMAIL_APP_PASSWORD"))
	var emails []string

	for _, contact := range campaing.Contacts {
		emails = append(emails, contact.Email)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("EMAIL_USER"))
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", campaing.Name)
	message.SetBody("text/html", campaing.Content)

	fmt.Println("Chegou aqui")
	return dialer.DialAndSend(message)
}
