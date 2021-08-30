package main

import (
	"context"
	"log"
	"net/http"
	"notification/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	proto.RegisterNotificationServiceHandlerFromEndpoint(context.Background(), mux, "localhost:4000", opts)

	log.Println("running grpc gateway on 8081")
	http.ListenAndServe(":8081", mux)
}

/*
endpoint: POST endpoint localhost:8081/proto3.NotificationService/SendEmail
body:
{
	"EmailRequests": [
		{
			"Email": "abc",
		"Subject": "aaaaaa",
		"Body": "aaaa"
		}
	]
}
*/
