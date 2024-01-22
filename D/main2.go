package main

import "fmt"

type MessageSender interface {
	Send(to string, message string)
}

type EmailService struct{}

func (e *EmailService) Send(to string, message string) {
	fmt.Printf("Sending Email to %s: %s\n", to, message)
}

type SMSService struct{}

func (s *SMSService) Send(to string, message string) {
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
}

type NotificationService struct {
	messageSender MessageSender
}

func (n *NotificationService) Notify(to string, message string) {
	n.messageSender.Send(to, message)
}

func main() {
	emailService := &EmailService{}
	smsService := &SMSService{}

	notificationServiceEmail := &NotificationService{messageSender: emailService}
	notificationServiceSMS := &NotificationService{messageSender: smsService}

	notificationServiceEmail.Notify("test@example.com", "Hello via Email!")
	notificationServiceSMS.Notify("1234567890", "Hello via SMS!")
}
