package services

import (
	"context"
	"log"
	"ng-book-service/helpers"
	"ng-book-service/proto"
	"ng-book-service/repositories"
)

type ReviewService struct{ proto.ReviewServiceServer }

func (s *ReviewService) Get(ctx context.Context, in *proto.Empty) (*proto.Reviews, error) {
	log.Println("[*] [ReviewService] Get")

	reviews, err := repositories.Review.Get()
	if err != nil {
		return nil, err
	}

	result := proto.Reviews{}

	for _, obj := range *reviews {
		review, err := helpers.Review.FromModelToProto(&obj)

		if err != nil {
			return nil, err
		}

		result.Reviews = append(result.Reviews, review)
	}

	return &result, nil
}

func (s *ReviewService) GetByBookId(ctx context.Context, in *proto.Id) (*proto.Reviews, error) {
	log.Println("[*] [ReviewService] Get by book id")

	reviews, err := repositories.Review.GetByBookId(in.Id)
	if err != nil {
		return nil, err
	}

	result := proto.Reviews{}

	for _, obj := range *reviews {
		review, err := helpers.Review.FromModelToProto(&obj)

		if err != nil {
			return nil, err
		}

		result.Reviews = append(result.Reviews, review)
	}

	return &result, nil
}

func (s *ReviewService) GetByAccountId(ctx context.Context, in *proto.Id) (*proto.Reviews, error) {
	log.Println("[*] [ReviewService] Get by account id")

	reviews, err := repositories.Review.GetByAccountId(in.Id)
	if err != nil {
		return nil, err
	}

	result := proto.Reviews{}

	for _, obj := range *reviews {
		review, err := helpers.Review.FromModelToProto(&obj)

		if err != nil {
			return nil, err
		}

		result.Reviews = append(result.Reviews, review)
	}

	return &result, nil
}

func (s *ReviewService) GetById(ctx context.Context, in *proto.Id) (*proto.Review, error) {
	log.Println("[*] [ReviewService] GetById")

	review, err := repositories.Review.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	result, err := helpers.Review.FromModelToProto(review)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ReviewService) Insert(ctx context.Context, in *proto.Review) (*proto.Status, error) {
	log.Println("[*] [ReviewService] Insert")

	review, err := helpers.Review.FromProtoToModel(in)

	if err != nil {
		return nil, err
	}

	_, err = repositories.Review.Insert(*review)

	if err != nil {
		return nil, err
	}

	repositories.Review.RecalculateRating(review.BookId.String())

	return &proto.Status{Code: 0, Message: "Ok"}, nil
}

// func (s *ReviewService) Update(ctx context.Context, in *proto.Review) (*proto.Status, error) {
// 	log.Println("[*] [ReviewService] Update")

// 	review, err := helpers.Review.FromProtoToModel(in)

// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = repositories.Review.Update(*review)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &proto.Status{Code: 0, Message: "Ok"}, nil
// }

func (s *ReviewService) Delete(ctx context.Context, in *proto.IdAccountId) (*proto.Status, error) {
	log.Println("[*] [ReviewService] Delete")

	if _, err := repositories.Review.Delete(in.Id, in.AccountId); err != nil {
		return nil, err
	}

	return &proto.Status{Code: 0, Message: "OK"}, nil
}
