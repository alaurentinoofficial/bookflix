package models

import (
	"ng-catalog-service/proto"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Catalog struct {
	Id        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	Items     []CatalogItem `json:"items" gorm:"preload:true"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"update_at,omitempty"`
}

func (this *Catalog) BeforeCreate(scope *gorm.Scope) error {
	uu, _ := uuid.NewV4()
	scope.SetColumn("Id", uu.String())

	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", nil)

	return nil
}

func (this *Catalog) FromProto(proto *proto.Catalog) {
	uuid_catalog, _ := uuid.FromString(proto.Id)

	this.Id = uuid_catalog
	this.Name = proto.Name
	this.Items = []CatalogItem{}

	for _, item := range proto.Items {
		i := CatalogItem{}
		i.FromProto(item)

		this.Items = append(this.Items, i)
	}
}

func (this *Catalog) Proto() *proto.Catalog {
	result := proto.Catalog{}

	result.Id = this.Id.String()
	result.Name = this.Name
	result.Items = []*proto.CatalogItem{}

	for _, item := range this.Items {
		result.Items = append(result.Items, item.Proto())
	}

	return &result
}
