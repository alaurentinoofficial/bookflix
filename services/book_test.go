package services

import (
	"context"
	"fmt"
	"ng-book-service/proto"
	"testing"

	"google.golang.org/grpc"
)

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
