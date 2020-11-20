package repositories

import (
	"ng-book-service/db"
	"ng-book-service/models"
)

type bookRepository struct{}

func (this *bookRepository) Get() (*[]models.Book, error) {
	result := []models.Book{}

	if err := db.Get().Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *bookRepository) GetById(id string) (*models.Book, error) {
	result := models.Book{}

	if err := db.Get().Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *bookRepository) Insert(book models.Book) (*models.Book, error) {
	if err := db.Get().Create(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (this *bookRepository) Update(book models.Book) (*models.Book, error) {
	var count int
	if err := db.Get().Model(models.Book{}).Where("id = ?", book.Id).Count(&count).Updates(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (this *bookRepository) Delete(id string) (*[]models.Book, error) {
	result := []models.Book{}

	if err := db.Get().Where("id = ?", id).Delete(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

var (
	Book bookRepository = bookRepository{}
)
