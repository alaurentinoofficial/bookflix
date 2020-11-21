package repositories

import (
	"ng-book-service/db"
	"ng-book-service/models"
)

type categoryRepository struct{}

func (this *categoryRepository) Get() (*[]models.Category, error) {
	result := []models.Category{}

	if err := db.Get().Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *categoryRepository) GetById(id string) (*models.Category, error) {
	result := models.Category{}

	if err := db.Get().Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *categoryRepository) Insert(category models.Category) (*models.Category, error) {
	if err := db.Get().Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (this *categoryRepository) Update(category models.Category) (*models.Category, error) {
	var count int
	if err := db.Get().Model(models.Category{}).Where("id = ?", category.Id).Count(&count).Updates(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (this *categoryRepository) Delete(id string) (*[]models.Category, error) {
	result := []models.Category{}

	if err := db.Get().Where("id = ?", id).Delete(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

var (
	Category categoryRepository = categoryRepository{}
)
