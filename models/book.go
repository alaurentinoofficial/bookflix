package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Book struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key"`
	Title     string    `json:"title,omitempty" validate:"required,min=4"`
	Author    string    `json:"author,omitempty" validate:"required,min=4"`
	Resume    string    `json:"resume,omitempty" validate:"required,min=20"`
	Rating    float32   `json:"rating,omitempty"`
	ImageUrl  string    `json:"image_url,omitempty" validate:"required,url"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`

	Categories []Category `gorm:"many2many:book_categories;"`
	Reviews    []Review   `gorm:"foreignKey:BookId;references:BookId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}

func (this *Book) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())

	scope.SetColumn("Rating", 0)

	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", nil)

	return nil
}

func (this *Book) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
