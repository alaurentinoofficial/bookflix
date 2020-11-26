package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BannerItem struct {
	Id          uuid.UUID `json:"id"`
	Placeholder string    `json:"placeholder,omitempty"`
	Link        string    `json:"link"`
	ImageUrl    string    `json:"image_url"`
}

type Banner struct {
	Id        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Items     []BannerItem `json:"items"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"update_at,omitempty"`
}

func (this *Banner) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())

	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", nil)

	return nil
}
