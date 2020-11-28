package services

import (
	"context"
	"log"
	"ng-auth-service/models"
	"ng-auth-service/proto"
	"ng-auth-service/repositories"
)

type AccountService struct{ proto.AccountServiceServer }

func (s *AccountService) Auth(ctx context.Context, in *proto.Login) (*proto.Status, error) {
	log.Println("[*] [AccountService] Auth")

	account, err := repositories.Account.Auth(in.Email, in.Password)
	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "OK", Id: account.Id.String()}, nil
}

func (s *AccountService) Get(ctx context.Context, in *proto.Empty) (*proto.ResponseAccounts, error) {
	log.Println("[*] [AccountService] Get")

	accounts, err := repositories.Account.Get()
	if err != nil {
		return nil, err
	}

	result := proto.ResponseAccounts{}

	for _, obj := range *accounts {
		result.Accounts = append(result.Accounts, obj.Proto())
	}

	return &result, nil
}

func (s *AccountService) GetById(ctx context.Context, in *proto.Id) (*proto.ResponseAccount, error) {
	log.Println("[*] [AccountService] GetById")

	account, err := repositories.Account.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	return account.Proto(), nil
}

func (s *AccountService) Insert(ctx context.Context, in *proto.CreateAccount) (*proto.Status, error) {
	log.Println("[*] [AccountService] Insert")

	account := models.Account{}
	account.FromProto(in)

	_, err := repositories.Account.Insert(account)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok", Id: account.Id.String()}, nil
}

func (s *AccountService) Delete(ctx context.Context, in *proto.Id) (*proto.Status, error) {
	log.Println("[*] [AccountService] Delete")

	_, err := repositories.Account.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &proto.Status{}, nil
}
