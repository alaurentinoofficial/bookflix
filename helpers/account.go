package helpers

import (
	"ng-auth-service/models"
	"ng-auth-service/proto"
)

type account struct{}

func (account) FromModelToProto(acc *models.Account) (*proto.ResponseAccount, error) {
	result := proto.ResponseAccount{
		Id:        acc.Id.String(),
		Name:      acc.Name,
		Email:     acc.Email,
		Profile:   acc.Profile,
		CreatedAt: acc.CreatedAt.String(),
		UpdatedAt: acc.UpdatedAt.String(),
	}

	return &result, nil
}

func (account) FromProtoToModel(acc *proto.CreateAccount) (*models.Account, error) {
	result := models.Account{
		Name:     acc.Name,
		Email:    acc.Email,
		Password: acc.Password,
	}

	return &result, nil
}

var Account account = account{}
