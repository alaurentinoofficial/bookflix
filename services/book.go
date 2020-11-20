package services

import (
	"context"
	"log"
	"ng-book-service/proto"
)

type BookService struct{ proto.BookServiceServer }

func (s *BookService) Get(ctx context.Context, in *proto.Empty) (*proto.Books, error) {
	log.Println("[*] Get")

	return &proto.Books{}, nil
}

func (s *BookService) GetById(ctx context.Context, in *proto.Id) (*proto.Book, error) {
	log.Println("[*] GetById")

	return &proto.Book{}, nil
}

func (s *BookService) Insert(ctx context.Context, in *proto.Book) (*proto.Status, error) {
	log.Println("[*] Insert")

	return &proto.Status{}, nil
}

func (s *BookService) Update(ctx context.Context, in *proto.Book) (*proto.Status, error) {
	log.Println("[*] Insert")

	return &proto.Status{}, nil
}

func (s *BookService) Delete(ctx context.Context, in *proto.Id) (*proto.Status, error) {
	log.Println("[*] Insert")

	return &proto.Status{}, nil
}
