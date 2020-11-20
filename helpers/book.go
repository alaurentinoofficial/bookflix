package helpers

import (
	"ng-book-service/models"
	"ng-book-service/proto"

	uuid "github.com/satori/go.uuid"
)

type book struct{}

func (book) FromModelToProto(book *models.Book) (*proto.Book, error) {
	categories := []*proto.Category{}

	for _, obj := range book.Categories {
		categories = append(categories, &proto.Category{
			Id:   obj.Id.String(),
			Name: obj.Name,
		})
	}

	result := proto.Book{
		Id:         book.Id.String(),
		Title:      book.Title,
		Author:     book.Author,
		Resume:     book.Resume,
		Rating:     book.Rating,
		ImageUrl:   book.ImageUrl,
		CreatedAt:  book.CreatedAt.String(),
		UpdatedAt:  book.UpdatedAt.String(),
		Categories: categories,
	}

	return &result, nil
}

func (book) FromProtoToModel(book *proto.Book) (*models.Book, error) {
	categories := []models.Category{}

	for _, obj := range book.Categories {
		uu, err := uuid.FromString(book.Id)

		if err != nil {
			return nil, err
		}

		categories = append(categories, models.Category{
			Id:   uu,
			Name: obj.Name,
		})
	}

	uu, err := uuid.FromString(book.Id)

	if err != nil {
		return nil, err
	}

	result := models.Book{
		Id:         uu,
		Title:      book.Title,
		Author:     book.Author,
		Resume:     book.Resume,
		Rating:     book.Rating,
		ImageUrl:   book.ImageUrl,
		Categories: categories,
	}

	return &result, nil
}

var Book book = book{}
