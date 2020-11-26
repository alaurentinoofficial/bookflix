package services

import (
	"context"
	"log"
	"ng-catalog-service/models"
	"ng-catalog-service/proto"
	"ng-catalog-service/repositories"
)

type BannerService struct{ proto.BannerServiceServer }

func (s *BannerService) Get(ctx context.Context, in *proto.Empty) (*proto.Banners, error) {
	log.Println("[*] [BannerService] Get")

	banners, err := repositories.Banner.Get()
	if err != nil {
		return nil, err
	}

	result := proto.Banners{}

	for _, obj := range *banners {
		result.Banners = append(result.Banners, obj.Proto())
	}

	return &result, nil
}

func (s *BannerService) GetById(ctx context.Context, in *proto.Id) (*proto.Banner, error) {
	log.Println("[*] [BannerService] GetById")

	catalog, err := repositories.Banner.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	result := catalog.Proto()

	return result, nil
}

func (s *BannerService) Insert(ctx context.Context, in *proto.Banner) (*proto.Status, error) {
	log.Println("[*] [BannerService] Insert")

	catalog := models.Banner{}
	catalog.FromProto(in)

	_, err := repositories.Banner.Insert(catalog)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

func (s *BannerService) Update(ctx context.Context, in *proto.Banner) (*proto.Status, error) {
	log.Println("[*] [BannerService] Update")

	catalog := models.Banner{}
	catalog.FromProto(in)

	_, err := repositories.Banner.Update(catalog)

	if err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

func (s *BannerService) Delete(ctx context.Context, in *proto.Id) (*proto.Status, error) {
	log.Println("[*] [BannerService] Delete")

	_, err := repositories.Banner.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &proto.Status{}, nil
}
