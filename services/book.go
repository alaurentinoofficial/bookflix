package services

import (
	"context"
	"log"
	"ng-book-service/helpers"
	"ng-book-service/proto"
	"ng-book-service/repositories"
)

type BookService struct{ proto.BookServiceServer }

func (s *BookService) Get(ctx context.Context, in *proto.Empty) (*proto.Books, error) {
	log.Println("[*] Get")

	books, err := repositories.Book.Get()
	if err != nil {
		return nil, err
	}

	result := proto.Books{}

	for _, obj := range *books {
		book, err := helpers.Book.FromModelToProto(&obj)

		if err != nil {
			return nil, err
		}

		result.Books = append(result.Books, book)
	}

	return &result, nil
}

func (s *BookService) GetById(ctx context.Context, in *proto.Id) (*proto.Book, error) {
	log.Println("[*] GetById")

	book, err := repositories.Book.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	result, err := helpers.Book.FromModelToProto(book)

	if err != nil {
		return nil, err
	}

	return result, nil
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
