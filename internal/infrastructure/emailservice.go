package infrastructure

import "fmt"

type SMTPEmailService struct {
}

func (e *SMTPEmailService) Send(to string, subject string, body string) error {
	fmt.Printf("Sending email to %s with subject %s and body %s\n", to, subject, body)
	return nil
}
