package services

import (
	"context"
	"fmt"
	"ng-auth-service/proto"
	"testing"

	"google.golang.org/grpc"
)

var (
	account = proto.ResponseAccount{}
)

func TestIntegrationMain(t *testing.T) {
	t.Run("insert", Insert)
	t.Run("login", Login)
	t.Run("get", Get)
	t.Run("get by id", GetById)
	t.Run("delete", Delete)
	t.Run("get deleted", GetDeleted)
}

func Insert(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9090"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewAccountServiceClient(conn)

	account_request := proto.CreateAccount{
		Name:     "Anderson Laurentino",
		Email:    "anderson@laurentino.me",
		Password: "1234567890n",
	}

	response, err := c.Insert(context.Background(), &account_request)
	if err != nil {
		t.Errorf(err.Error())
	}

	if response.Code != 0 {
		t.Errorf("Not succefull")
	}
}

func Get(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9090"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewAccountServiceClient(conn)

	response, err := c.Get(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(response.Accounts) == 0 {
		t.Errorf("cateories is void")
	}

	account = *response.Accounts[0]
}

func GetDeleted(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9090"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewAccountServiceClient(conn)

	response, err := c.Get(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(response.Accounts) != 0 {
		t.Errorf("cateories is void")
	}
}

func Login(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9090"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewAccountServiceClient(conn)

	login_request := proto.Login{
		Email:    "anderson@laurentino.me",
		Password: "1234567890n",
	}

	response, err := c.Auth(context.Background(), &login_request)
	if err != nil {
		t.Errorf(err.Error())
	}

	if response.Code != 0 {
		t.Errorf("Not succefull")
	}
}

func GetById(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9090"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewAccountServiceClient(conn)

	response, err := c.GetById(context.Background(), &proto.Id{Id: account.Id})
	if err != nil {
		t.Errorf(err.Error())
	}

	if !(response.Email == "anderson@laurentino.me" && response.Name == "Anderson Laurentino") {
		t.Log(response)
		t.Errorf("account not found")
	}
}

func Delete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9090"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewAccountServiceClient(conn)

	response, err := c.Delete(context.Background(), &proto.Id{Id: account.Id})
	if err != nil {
		t.Errorf(err.Error())
	}

	if response.Code != 0 {
		t.Errorf("Not succefull")
	}
}
