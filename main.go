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

	// category := models.Category{Name: "Romance"}

	// db.Get().Create(&category)

	// book := models.Book{
	// 	Title:      "Como Mentir com Estatística",
	// 	Author:     "Darrel Huff",
	// 	Resume:     "Como Mentir com Estatística",
	// 	Rating:     4,
	// 	ImageUrl:   "http://www.exemplo.com",
	// 	Categories: []models.Category{models.Category{Id: category.Id}},
	// }

	// repositories.Book.Insert(book)

	// fmt.Println(book)

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
