package main

import (
	"client/proto"
	context "context"
	"fmt"
	"log"

	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	emailService := proto.NewNotificationServiceClient(conn)
	sendSingleEmail(emailService)
	sendBulkEmail(emailService)
}

func sendSingleEmail(emailService proto.NotificationServiceClient) {
	req := proto.EmailRequest{
		Subject: "haloo",
		Body:    "test",
		Email:   "test@gmail.com",
	}
	result, err := emailService.SendEmail(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result is", result.Message)
}

func sendBulkEmail(emailService proto.NotificationServiceClient) {
	var emails []*proto.EmailRequest

	for i := 0; i < 10; i++ {
		emails = append(emails, &proto.EmailRequest{
			Subject: "haloo",
			Body:    "test",
			Email:   fmt.Sprintf("test%d@gmail.com", i),
		})
	}
	req := proto.EmailBulkRequest{
		EmailRequests: emails,
	}
	result, err := emailService.SendBulkEmail(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result for send bulk email is", result.Message)
}
