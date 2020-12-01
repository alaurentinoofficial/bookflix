package models

import (
	"ng-catalog-service/proto"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CatalogItem struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"Name"`
	CatalogId uuid.UUID `json:"catalog_id"`
	ItemId    uuid.UUID `json:"item_id"`
	ImageUrl  string    `json:"image_url,omitempty"`
}

func (this *CatalogItem) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())

	return nil
}

func (this *CatalogItem) FromProto(proto *proto.CatalogItem) {
	uu, _ := uuid.FromString(proto.Id)
	uuid_item, _ := uuid.FromString(proto.ItemId)

	this.Id = uu
	this.ItemId = uuid_item
	this.Name = proto.Name
	this.ImageUrl = proto.ImageUrl
}

func (this *CatalogItem) Proto() *proto.CatalogItem {
	return &proto.CatalogItem{
		Id:       this.Id.String(),
		ItemId:   this.ItemId.String(),
		Name:     this.Name,
		ImageUrl: this.ImageUrl,
	}
}
