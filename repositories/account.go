package repositories

import (
	"ng-auth-service/db"
	"ng-auth-service/models"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type accountRepository struct{}

func (this *accountRepository) Auth(email string, password string) (*models.Account, error) {
	result := models.Account{}

	if err := db.Get().Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}

	errPwd := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if errPwd != nil && errPwd == bcrypt.ErrMismatchedHashAndPassword {
		return nil, status.Error(codes.InvalidArgument, "Email or Password invalid!")
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
	count := 0

	if err := db.Get().Table("accounts").Where("email = ?", account.Email).Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, status.Error(codes.AlreadyExists, "Email already exists!")
	}

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
