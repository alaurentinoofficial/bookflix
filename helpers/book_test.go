package helpers

import (
	"ng-book-service/models"
	"ng-book-service/proto"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestBookHelper(t *testing.T) {
	t.Run("from model to proto", BookFromModelToProto)
	t.Run("from proto to model", BookFromProtoToModel)
}

func BookFromModelToProto(t *testing.T) {
	bookId, _ := uuid.FromString("05cb0e6f-d4b0-407c-92c2-f3c1d9f85958")
	categoryId, _ := uuid.FromString("72a7cbdf-5780-41e5-8bc9-3296cef52a8c")
	createAt, _ := time.Parse("2006-01-02 15:04", "2020-11-21 12:38")
	updateAt, _ := time.Parse("2006-01-02 15:04", "2020-11-21 14:38")

	model := models.Book{
		Id:        bookId,
		Title:     "Introducing Go: Build Reliable, Scalable Programs",
		Author:    "Caleb Doxsey",
		Resume:    "Perfect for beginners familiar with programming basics ...",
		Rating:    3.7,
		ImageUrl:  "http://www.contoso.com/static/kdjfldksajffakdsjfl.png",
		CreatedAt: createAt,
		UpdatedAt: updateAt,
		Categories: []models.Category{
			models.Category{
				Id:   categoryId,
				Name: "Programming",
			},
		},
	}

	proto, err := Book.FromModelToProto(&model)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !(proto.Id == "05cb0e6f-d4b0-407c-92c2-f3c1d9f85958" &&
		proto.Title == "Introducing Go: Build Reliable, Scalable Programs" &&
		proto.Author == "Caleb Doxsey" &&
		proto.Resume == "Perfect for beginners familiar with programming basics ..." &&
		proto.Rating == 3.7 &&
		proto.ImageUrl == "http://www.contoso.com/static/kdjfldksajffakdsjfl.png" &&
		proto.CreatedAt == "2020-11-21 12:38:00 +0000 UTC" &&
		proto.UpdatedAt == "2020-11-21 14:38:00 +0000 UTC" &&
		len(proto.Categories) == 1 &&
		proto.Categories[0].Id == "72a7cbdf-5780-41e5-8bc9-3296cef52a8c" &&
		proto.Categories[0].Name == "Programming") {

		t.Errorf("result objet is not ok")
	}
}

func BookFromProtoToModel(t *testing.T) {
	proto := proto.Book{
		Id:        "05cb0e6f-d4b0-407c-92c2-f3c1d9f85958",
		Title:     "Introducing Go: Build Reliable, Scalable Programs",
		Author:    "Caleb Doxsey",
		Resume:    "Perfect for beginners familiar with programming basics ...",
		Rating:    3.7,
		ImageUrl:  "http://www.contoso.com/static/kdjfldksajffakdsjfl.png",
		CreatedAt: "2020-11-21 12:38:00 +0000 UTC",
		UpdatedAt: "2020-11-21 14:38:00 +0000 UTC",
		Categories: []*proto.Category{
			&proto.Category{
				Id:   "72a7cbdf-5780-41e5-8bc9-3296cef52a8c",
				Name: "Programming",
			},
		},
	}

	model, err := Book.FromProtoToModel(&proto)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !(model.Id.String() == "05cb0e6f-d4b0-407c-92c2-f3c1d9f85958" &&
		model.Title == "Introducing Go: Build Reliable, Scalable Programs" &&
		model.Author == "Caleb Doxsey" &&
		model.Resume == "Perfect for beginners familiar with programming basics ..." &&
		model.Rating == 3.7 &&
		model.ImageUrl == "http://www.contoso.com/static/kdjfldksajffakdsjfl.png" &&
		len(model.Categories) == 1 &&
		model.Categories[0].Id.String() == "72a7cbdf-5780-41e5-8bc9-3296cef52a8c" &&
		model.Categories[0].Name == "Programming") {

		t.Log(model.Id.String() == "05cb0e6f-d4b0-407c-92c2-f3c1d9f85958")
		t.Log(model.Categories[0].Id.String() == "72a7cbdf-5780-41e5-8bc9-3296cef52a8c")
		t.Log(model)
		t.Errorf("result objet is not ok")
	}
}
