package main

import (
	"fmt"
	"log"
	"net"
	"ng-auth-service/db"
	"ng-auth-service/proto"
	"ng-auth-service/services"
	"os"

	"google.golang.org/grpc"
)

var (
	port = ":9090"
)

func main() {
	db.Init()

	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterAccountServiceServer(s, &services.AccountService{})

	log.Println("[*] Listening in", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
