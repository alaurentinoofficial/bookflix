package services

import (
	"context"
	"fmt"
	"ng-book-service/proto"
	"testing"

	"google.golang.org/grpc"
)

var (
	book = proto.Book{}
)

func TestInsert(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	// cc := []*proto.Category{&proto.Category{Id: "1776209b-3332-42cd-94f3-480121d595ff", Name: "asdf"}}

	book := proto.Book{
		Title:    "Como Mentir com Estatística",
		Author:   "Darrel Huff",
		Resume:   "Como Mentir com Estatística",
		Rating:   4,
		ImageUrl: "http://www.exemplo.com",
		// Categories: cc,
	}

	fmt.Println(book)
	response, err := c.Insert(context.Background(), &book)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(response)
}

func TestGet(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	response, err := c.Get(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	book = *response.Books[0]
}

func TestUpdate(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	book.Title = "The drunkard's walk"

	response, err := c.Update(context.Background(), &proto.Book{Id: book.Id, Title: "The drunkard's walk"})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	fmt.Println(response)
}

func TestDelete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	response, err := c.Delete(context.Background(), &proto.Id{Id: book.Id})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	fmt.Println(response)
}
