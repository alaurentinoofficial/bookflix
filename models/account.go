package models

import (
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

	return nil
}

func (this *Account) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
