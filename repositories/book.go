package repositories

import (
	"ng-auth-service/db"
	"ng-auth-service/models"
)

type accountRepository struct{}

func (this *accountRepository) Auth(email string, password string) (*models.Account, error) {
	result := models.Account{}

	if err := db.Get().Where("email = ? and password = ?", email, password).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *accountRepository) Get() (*[]models.Account, error) {
	result := []models.Account{}

	if err := db.Get().Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *accountRepository) GetById(id string) (*models.Account, error) {
	result := models.Account{}

	if err := db.Get().Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *accountRepository) Insert(account models.Account) (*models.Account, error) {
	if err := db.Get().Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

// func (this *accountRepository) Update(account models.Account) (*models.Account, error) {
// 	var count int
// 	if err := db.Get().Model(models.Account{}).Where("id = ?", account.Id).Count(&count).Updates(&account).Error; err != nil {
// 		return nil, err
// 	}

// 	return &account, nil
// }

func (this *accountRepository) Delete(id string) (*[]models.Account, error) {
	result := []models.Account{}

	if err := db.Get().Where("id = ?", id).Delete(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

var (
	Account accountRepository = accountRepository{}
)
