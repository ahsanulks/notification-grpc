package service

import (
	"notification/model"
	"notification/proto"
)

// EmailService will provide all about email
type EmailService struct {
	emailRepository EmailRepository
	proto.UnimplementedNotificationServiceServer
}

// EmailRepository will provide logic to all email service needed
type EmailRepository interface {
	SendEmail(subject, body, to string) error
	SendBulkEmail([]model.Email) error
}

// NewEmailService create new email repository
func NewEmailService(emailRepository EmailRepository) *EmailService {
	return &EmailService{emailRepository, proto.UnimplementedNotificationServiceServer{}}
}

var (
	ErrorSendEmail   = "something when wrong when send email"
	SuccessSendEmail = "successfully send email"
)
