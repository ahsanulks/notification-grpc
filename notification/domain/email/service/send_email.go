package service

import (
	"context"
	"notification/proto"
)

func (es EmailService) SendEmail(ctx context.Context, request *proto.EmailRequest) (*proto.EmailReply, error) {
	err := es.emailRepository.SendEmail(request.Subject, request.Body, request.Email)
	if err != nil {
		return &proto.EmailReply{Message: ErrorSendEmail}, err
	}
	return &proto.EmailReply{Message: SuccessSendEmail}, nil
}
