package helpers

import (
	"ng-book-service/models"
	"ng-book-service/proto"

	uuid "github.com/satori/go.uuid"
)

type review struct{}

func (review) FromModelToProto(category *models.Review) (*proto.Review, error) {
	result := proto.Review{
		Id:           category.Id.String(),
		BookId:       category.BookId.String(),
		AccountId:    category.AccountId.String(),
		Rating:       category.Rating,
		PosNegRating: category.PosNegRating,
		Resume:       category.Resume,
		CreatedAt:    category.CreatedAt.String(),
		UpdatedAt:    category.UpdatedAt.String(),
	}

	return &result, nil
}

func (review) FromProtoToModel(category *proto.Review) (*models.Review, error) {
	reviewId, _ := uuid.FromString(category.Id)
	bookId, _ := uuid.FromString(category.BookId)
	accountId, _ := uuid.FromString(category.AccountId)

	result := models.Review{
		Id:           reviewId,
		BookId:       bookId,
		AccountId:    accountId,
		Resume:       category.Resume,
		Rating:       category.Rating,
		PosNegRating: category.PosNegRating,
	}

	return &result, nil
}

var Review review = review{}
