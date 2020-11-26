package services

import (
	"context"
	"log"
	"ng-catalog-service/models"
	"ng-catalog-service/proto"
	"ng-catalog-service/repositories"
)

type CatalogService struct{ proto.CatalogServiceServer }

func (s *CatalogService) Get(ctx context.Context, in *proto.Empty) (*proto.Catalogs, error) {
	log.Println("[*] [CatalogService] Get")

	catalogs, err := repositories.Catalog.Get()
	if err != nil {
		return nil, err
	}

	result := proto.Catalogs{}

	for _, obj := range *catalogs {
		result.Catalogs = append(result.Catalogs, obj.Proto())
	}

	return &result, nil
}

func (s *CatalogService) GetById(ctx context.Context, in *proto.Id) (*proto.Catalog, error) {
	log.Println("[*] [CatalogService] GetById")

	catalog, err := repositories.Catalog.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	result := catalog.Proto()

	return result, nil
}

func (s *CatalogService) Insert(ctx context.Context, in *proto.Catalog) (*proto.Status, error) {
	log.Println("[*] [CatalogService] Insert")

	catalog := models.Catalog{}
	catalog.FromProto(in)

	_, err := repositories.Catalog.Insert(catalog)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

func (s *CatalogService) Update(ctx context.Context, in *proto.Catalog) (*proto.Status, error) {
	log.Println("[*] [CatalogService] Update")

	catalog := models.Catalog{}
	catalog.FromProto(in)

	_, err := repositories.Catalog.Update(catalog)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

func (s *CatalogService) Delete(ctx context.Context, in *proto.Id) (*proto.Status, error) {
	log.Println("[*] [CatalogService] Delete")

	_, err := repositories.Catalog.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &proto.Status{}, nil
}
