package models

import (
	"ng-catalog-service/proto"
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

func (this *BannerItem) FromProto(proto *proto.BannerItem) {
	uuid_item, _ := uuid.FromString(proto.Id)
	this.Id = uuid_item
	this.Placeholder = proto.Placeholder
	this.Link = proto.Link
	this.ImageUrl = proto.ImageUrl
}

func (this *BannerItem) Proto() *proto.BannerItem {
	return &proto.BannerItem{
		Id:          this.Id.String(),
		Placeholder: this.Placeholder,
		Link:        this.Link,
		ImageUrl:    this.ImageUrl,
	}
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

func (this *Banner) FromProto(proto *proto.Banner) {
	uuid_banner, _ := uuid.FromString(proto.Id)

	this.Id = uuid_banner
	this.Name = proto.Name
	this.Items = []BannerItem{}

	for _, item := range proto.Items {
		i := BannerItem{}
		i.FromProto(item)

		this.Items = append(this.Items, i)
	}
}

func (this *Banner) Proto() *proto.Banner {
	result := proto.Banner{}

	result.Id = this.Id.String()
	result.Name = this.Name
	result.Items = []*proto.BannerItem{}

	for _, item := range this.Items {
		result.Items = append(result.Items, item.Proto())
	}

	return &result
}
