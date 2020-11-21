package helpers

import (
	"ng-book-service/models"
	"ng-book-service/proto"

	uuid "github.com/satori/go.uuid"
)

type category struct{}

func (category) FromModelToProto(category *models.Category) (*proto.Category, error) {
	result := proto.Category{
		Id:        category.Id.String(),
		Name:      category.Name,
		CreatedAt: category.CreatedAt.String(),
		UpdatedAt: category.UpdatedAt.String(),
	}

	return &result, nil
}

func (category) FromProtoToModel(category *proto.Category) (*models.Category, error) {
	uu, _ := uuid.FromString(category.Id)

	result := models.Category{
		Id:   uu,
		Name: category.Name,
	}

	return &result, nil
}

var Category category = category{}
