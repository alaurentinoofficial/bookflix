package main

import (
	"fmt"
	"log"
	"net"
	"ng-catalog-service/db"
	"ng-catalog-service/proto"
	"ng-catalog-service/services"
	"os"

	"google.golang.org/grpc"
)

var (
	port = ":9000"
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
	proto.RegisterCatalogServiceServer(s, &services.CatalogService{})
	proto.RegisterBannerServiceServer(s, &services.BannerService{})

	log.Println("[*] Listening in", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
