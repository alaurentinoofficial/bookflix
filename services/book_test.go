package services

import (
	"context"
	"fmt"
	"ng-book-service/proto"
	"testing"

	"google.golang.org/grpc"
)

func TestInsert(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":8080"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	book := proto.Book{
		Title:    "Como Mentir com Estatística",
		Author:   "Darrel Huff",
		Resume:   "Como Mentir com Estatística",
		Rating:   4,
		ImageUrl: "http://www.exemplo.com",
	}

	response, err := c.Insert(context.Background(), &book)
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	fmt.Println(response)
}

func TestGet(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":8080"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	response, err := c.Get(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	fmt.Println(response.Books)
}
