package repositories

import (
	"ng-catalog-service/db"
	"ng-catalog-service/models"
)

type bannerRepository struct{}

func (this *bannerRepository) Get() (*[]models.Banner, error) {
	result := []models.Banner{}

	if err := db.Get().Preload("BannerItems").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *bannerRepository) GetById(id string) (*models.Banner, error) {
	result := models.Banner{}

	if err := db.Get().Preload("BannerItems").Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *bannerRepository) Insert(catalog models.Banner) (*models.Banner, error) {
	items := catalog.Items
	catalog.Items = []models.BannerItem{}

	if err := db.Get().Create(&catalog).Error; err != nil {
		return nil, err
	}

	if err := db.Get().Model(&catalog).Association("BannerItems").Replace(items).Error; err != nil {
		this.Delete(catalog.Id.String())
		return nil, err
	}

	return &catalog, nil
}

func (this *bannerRepository) Update(catalog models.Banner) (*models.Banner, error) {
	var count int
	if err := db.Get().Model(models.Banner{}).Where("id = ?", catalog.Id).Count(&count).Updates(&catalog).Error; err != nil {
		return nil, err
	}

	return &catalog, nil
}

func (this *bannerRepository) Delete(id string) (*[]models.Banner, error) {
	result := []models.Banner{}

	if err := db.Get().Where("id = ?", id).Delete(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

var (
	Banner bannerRepository = bannerRepository{}
)
