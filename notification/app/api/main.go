package main

import (
	"log"
	"net"
	er "notification/domain/email/repository"
	es "notification/domain/email/service"
	pr "notification/domain/partner/repository"
	"notification/proto"

	"google.golang.org/grpc"
)

func main() {
	partnerRepository := pr.NewPartnerRepository()
	emailRepository := er.NewEmailRepository(partnerRepository)
	emailService := es.NewEmailService(emailRepository)

	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	// create new grpc server
	srv := grpc.NewServer()

	proto.RegisterNotificationServiceServer(srv, emailService)
	log.Println("server will started at port 4000")
	// running grpc server over tcp at localhost:4000
	if err := srv.Serve(lis); err != nil {
		log.Println("error when running server because", err.Error())
		panic(err)
	}
}
