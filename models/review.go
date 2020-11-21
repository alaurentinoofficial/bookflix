package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Review struct {
	Id           uuid.UUID `json:"id,omitempty" gorm:"primary_key"`
	Rating       float32   `json:"rating,omitempty"`
	PosNegRating float32   `json:"pos_neg_rating,omitempty"`
	Resume       string    `json:"resume,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"update_at,omitempty"`

	BookId uuid.UUID `json:"book_id" gorm:"not null"`

	AccountId uuid.UUID `json:"account_id" gorm:"not null"`
}

func (review *Review) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())

	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", nil)

	return nil
}

func (this *Review) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
