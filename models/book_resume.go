package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BookResume struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"Name"`
	ImageUrl string    `json:"image_url,omitempty"`
}

func (this *BookResume) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())

	return nil
}
