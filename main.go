package main

import (
	"fmt"
	"log"
	"net"
	"ng-book-service/db"
	"ng-book-service/proto"
	"ng-book-service/services"
	"os"

	"google.golang.org/grpc"
)

var (
	port = ":50051"
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
	proto.RegisterBookServiceServer(s, &services.BookService{})

	log.Println("[*] Listening in", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
