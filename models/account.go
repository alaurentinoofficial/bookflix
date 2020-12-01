package models

import (
	"ng-auth-service/proto"
	"ng-auth-service/utils"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,passwd"`
	Profile   string    `json:"profile,omitempty" validate:"required,profile"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
}

func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())

	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", nil)
	scope.SetColumn("Profile", "NORMAL")

	scope.SetColumn("Password", utils.HashPassword(account.Password))

	return nil
}

func (this *Account) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (this *Account) Proto() *proto.ResponseAccount {
	result := proto.ResponseAccount{
		Id:        this.Id.String(),
		Name:      this.Name,
		Email:     this.Email,
		Profile:   this.Profile,
		CreatedAt: this.CreatedAt.String(),
		UpdatedAt: this.UpdatedAt.String(),
	}

	return &result
}

func (this *Account) FromProto(acc *proto.CreateAccount) {
	this.Name = acc.Name
	this.Email = acc.Email
	this.Password = acc.Password
}
