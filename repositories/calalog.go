package repositories

import (
	"ng-catalog-service/db"
	"ng-catalog-service/models"
)

type catalogRepository struct{}

func (this *catalogRepository) Get() (*[]models.Catalog, error) {
	result := []models.Catalog{}

	if err := db.Get().Preload("CatalogItems").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *catalogRepository) GetById(id string) (*models.Catalog, error) {
	result := models.Catalog{}

	if err := db.Get().Preload("CatalogItems").Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *catalogRepository) Insert(catalog models.Catalog) (*models.Catalog, error) {
	items := catalog.Items
	catalog.Items = []models.CatalogItem{}

	if err := db.Get().Create(&catalog).Error; err != nil {
		return nil, err
	}

	if err := db.Get().Model(&catalog).Association("CatalogItems").Replace(items).Error; err != nil {
		this.Delete(catalog.Id.String())
		return nil, err
	}

	return &catalog, nil
}

func (this *catalogRepository) Update(catalog models.Catalog) (*models.Catalog, error) {
	var count int
	if err := db.Get().Model(models.Catalog{}).Where("id = ?", catalog.Id).Count(&count).Updates(&catalog).Error; err != nil {
		return nil, err
	}

	return &catalog, nil
}

func (this *catalogRepository) Delete(id string) (*[]models.Catalog, error) {
	result := []models.Catalog{}

	if err := db.Get().Where("id = ?", id).Delete(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

var (
	Catalog catalogRepository = catalogRepository{}
)
