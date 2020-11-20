package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Id   uuid.UUID `json:"id,omitempty" gorm:"primary_key"`
	Name string    `json:"name,omitempty" gorm:"unique" validate:"required,min=2"`
}

func (this *Category) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())
	return nil
}
