package services

import (
	"context"
	"log"
	"ng-book-service/helpers"
	"ng-book-service/proto"
	"ng-book-service/repositories"
)

type CategoryService struct{ proto.CategoryServiceServer }

func (s *CategoryService) Get(ctx context.Context, in *proto.Empty) (*proto.Categories, error) {
	log.Println("[*] [CategoryService] Get")

	categories, err := repositories.Category.Get()
	if err != nil {
		return nil, err
	}

	result := proto.Categories{}

	for _, obj := range *categories {
		category, err := helpers.Category.FromModelToProto(&obj)

		if err != nil {
			return nil, err
		}

		result.Categories = append(result.Categories, category)
	}

	return &result, nil
}

func (s *CategoryService) GetById(ctx context.Context, in *proto.Id) (*proto.Category, error) {
	log.Println("[*] [CategoryService] GetById")

	category, err := repositories.Category.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	result, err := helpers.Category.FromModelToProto(category)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CategoryService) Insert(ctx context.Context, in *proto.Category) (*proto.Status, error) {
	log.Println("[*] [CategoryService] Insert")

	category, err := helpers.Category.FromProtoToModel(in)

	if err != nil {
		return nil, err
	}

	_, err = repositories.Category.Insert(*category)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

func (s *CategoryService) Update(ctx context.Context, in *proto.Category) (*proto.Status, error) {
	log.Println("[*] [CategoryService] Update")

	category, err := helpers.Category.FromProtoToModel(in)

	if err != nil {
		return nil, err
	}

	_, err = repositories.Category.Update(*category)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

func (s *CategoryService) Delete(ctx context.Context, in *proto.Id) (*proto.Status, error) {
	log.Println("[*] [CategoryService] Delete")

	if _, err := repositories.Category.Delete(in.Id); err != nil {
		return nil, err
	}

	return &proto.Status{}, nil
}
