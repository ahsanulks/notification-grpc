package service

import (
	"context"
	"notification/model"
	"notification/proto"
)

// SendBulkEmail will send bulk email at same time
func (es EmailService) SendBulkEmail(ctx context.Context, request *proto.EmailBulkRequest) (*proto.EmailBulkReply, error) {
	var (
		emails []model.Email
	)

	for _, emailRequest := range request.EmailRequests {
		emails = append(emails, model.Email{
			Subject: emailRequest.Subject,
			Body:    emailRequest.Body,
			To:      emailRequest.Email,
		})
	}
	err := es.emailRepository.SendBulkEmail(emails)
	if err != nil {
		return &proto.EmailBulkReply{Message: ErrorSendEmail}, err
	}
	return &proto.EmailBulkReply{Message: SuccessSendEmail}, nil
}
