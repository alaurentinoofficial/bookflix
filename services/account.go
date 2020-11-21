package services

import (
	"context"
	"log"
	"ng-auth-service/helpers"
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

	_, err = helpers.Account.FromModelToProto(account)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "OK"}, nil
}

func (s *AccountService) GetById(ctx context.Context, in *proto.Id) (*proto.ResponseAccount, error) {
	log.Println("[*] [AccountService] GetById")

	account, err := repositories.Account.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	result, err := helpers.Account.FromModelToProto(account)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *AccountService) Insert(ctx context.Context, in *proto.CreateAccount) (*proto.Status, error) {
	log.Println("[*] [AccountService] Insert")

	book, err := helpers.Account.FromProtoToModel(in)

	if err != nil {
		return nil, err
	}

	_, err = repositories.Account.Insert(*book)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

func (s *AccountService) Delete(ctx context.Context, in *proto.Id) (*proto.Status, error) {
	log.Println("[*] [AccountService] Delete")

	_, err := repositories.Account.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &proto.Status{}, nil
}
