package repositories

import (
	"ng-book-service/db"
	"ng-book-service/models"
)

type reviewRepository struct{}

func (this *reviewRepository) Get() (*[]models.Review, error) {
	result := []models.Review{}

	if err := db.Get().Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *reviewRepository) GetByAccountId(id string) (*[]models.Review, error) {
	result := []models.Review{}

	if err := db.Get().Where("account_id = ?", id).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *reviewRepository) GetByBookId(id string) (*[]models.Review, error) {
	result := []models.Review{}

	if err := db.Get().Where("book_id = ?", id).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *reviewRepository) GetById(id string) (*models.Review, error) {
	result := models.Review{}

	if err := db.Get().Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *reviewRepository) Insert(category models.Review) (*models.Review, error) {
	if err := db.Get().Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (this *reviewRepository) Update(category models.Review) (*models.Review, error) {
	var count int
	if err := db.Get().Model(models.Review{}).Where("id = ?", category.Id).Count(&count).Updates(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (this *reviewRepository) Delete(id string) (*[]models.Review, error) {
	result := []models.Review{}

	if err := db.Get().Where("id = ?", id).Delete(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

var (
	Review reviewRepository = reviewRepository{}
)
