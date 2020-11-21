package services

import (
	"context"
	"fmt"
	"ng-book-service/proto"
	"testing"

	"google.golang.org/grpc"
)

var (
	book            = proto.Book{}
	category        = proto.Category{}
	reviews_request = []proto.Review{}
	reviews         = proto.Reviews{}
)

func TestIntegrationMain(t *testing.T) {
	t.Run("category insert", CategoryInsert)
	t.Run("category get", CategoryGet)
	t.Run("category update", CategoryUpdate)
	t.Run("category get by id", CategoryGetById)

	t.Run("book insert", BookInsert)
	t.Run("book get", BookGet)
	t.Run("book update", BookUpdate)
	t.Run("book get by id", BookGetById)

	t.Run("review insert", ReviewInsert)
	t.Run("review get", ReviewGet)
	t.Run("review get by account id", ReviewGetByAccountId)
	t.Run("review delete", ReviewDelete)
	t.Run("review get by book id", ReviewGetByBookId)
	t.Run("review delete total", ReviewDeleteTotal)

	t.Run("book delete", BookDelete)
	t.Run("category delete", CategoryDelete)
}

func CategoryInsert(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCategoryServiceClient(conn)

	category_request := proto.Category{
		Name: "Romance",
	}

	response, err := c.Insert(context.Background(), &category_request)
	if err != nil {
		t.Errorf(err.Error())
	}

	if response.Code != 0 {
		t.Errorf("Not succefull")
	}
}

func CategoryGet(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCategoryServiceClient(conn)

	response, err := c.Get(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(response.Categories) == 0 {
		t.Errorf("cateories is void")
	}

	category = *response.Categories[0]
}

func CategoryUpdate(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCategoryServiceClient(conn)

	category_request := proto.Category{
		Id:   category.Id,
		Name: "Statistics",
	}

	response, err := c.Insert(context.Background(), &category_request)
	if err != nil {
		t.Errorf(err.Error())
	}

	if response.Code != 0 {
		t.Errorf("Not succefull")
	}
}

func CategoryGetById(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCategoryServiceClient(conn)

	response, err := c.GetById(context.Background(), &proto.Id{Id: category.Id})
	if err != nil {
		t.Errorf(err.Error())
	}

	if response.Name == "Statistics" {
		t.Errorf("category wasn't changed")
	}
}

func CategoryDelete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCategoryServiceClient(conn)

	response, err := c.Delete(context.Background(), &proto.Id{Id: category.Id})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	if response.Code != 0 {
		t.Errorf("not was deleted")
	}
}

func BookInsert(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	cc := []*proto.Category{&proto.Category{Id: category.Id}}

	book := proto.Book{
		Title:      "Como Mentir com Estatística",
		Author:     "Darrel Huff",
		Resume:     "Como Mentir com Estatística",
		Rating:     4,
		ImageUrl:   "http://www.exemplo.com",
		Categories: cc,
	}

	response, err := c.Insert(context.Background(), &book)
	if err != nil {
		fmt.Println(response)
		t.Errorf(err.Error())
	}

	if response.Code != 0 {
		fmt.Println(response)
		t.Errorf("Not succefull")
	}
}

func BookGet(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	response, err := c.Get(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	if len(response.Books) == 0 {
		t.Errorf("books is void")
	}

	book = *response.Books[0]
}

func BookUpdate(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	book.Title = "The drunkard's walk"

	response, err := c.Update(context.Background(), &proto.Book{Id: book.Id, Title: "The drunkard's walk"})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	if response.Code != 0 {
		t.Errorf("Not succefull")
	}
}

func BookGetById(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	response, err := c.GetById(context.Background(), &proto.Id{Id: book.Id})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	if response.Title != "The drunkard's walk" {
		t.Errorf("book wasn't changed")
	}
}

func BookDelete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBookServiceClient(conn)

	response, err := c.Delete(context.Background(), &proto.Id{Id: book.Id})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	if response.Code != 0 {
		t.Errorf("not was deleted")
	}
}

func ReviewInsert(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewReviewServiceClient(conn)

	reviews_request = []proto.Review{
		proto.Review{
			Resume:    "Really awesome",
			Rating:    4.7,
			BookId:    book.Id,
			AccountId: "5ad21183-a43e-48dd-ad91-10d7c5c681dc",
		},
		proto.Review{
			Resume:    "I loved, explain very well the randomness in our lifes",
			Rating:    5,
			BookId:    book.Id,
			AccountId: "42653ae4-23dd-46c6-93e5-681c56f039fb",
		},
		proto.Review{
			Resume:    "I hated... unnecessary",
			Rating:    1.2,
			BookId:    book.Id,
			AccountId: "e25efb20-dc56-4bb2-bdd3-ccb1ca01c892",
		},
	}

	for _, review := range reviews_request {
		response, err := c.Insert(context.Background(), &review)
		if err != nil {
			t.Errorf(err.Error())
		}

		if response.Code != 0 {
			t.Errorf("Not succefull")
		}
	}
}

func ReviewGet(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewReviewServiceClient(conn)

	response, err := c.Get(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(response.Reviews) != 3 {
		t.Errorf("reviews hasn't the correct size")
	}

	for _, o := range reviews_request {
		result := false

		for _, r := range response.Reviews {
			if o.Resume == r.Resume &&
				o.Rating == r.Rating &&
				o.AccountId == r.AccountId &&
				o.BookId == r.BookId {

				result = true
				break
			}
		}

		if !result {
			t.Errorf("review not found")
		}
	}

	reviews = *response
}

func ReviewGetById(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewReviewServiceClient(conn)

	response, err := c.GetById(context.Background(), &proto.Id{Id: reviews.Reviews[0].Id})
	if err != nil {
		t.Errorf(err.Error())
	}

	result := false
	for _, o := range reviews_request {
		if o.Id == response.Id {
			result = true
			break
		}
	}

	if !result {
		t.Log(response)
		t.Errorf("review not found")
	}
}

func ReviewGetByAccountId(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewReviewServiceClient(conn)

	response, err := c.GetByAccountId(context.Background(), &proto.Id{Id: "5ad21183-a43e-48dd-ad91-10d7c5c681dc"})
	if err != nil {
		fmt.Println(response)
		t.Errorf(err.Error())
	}

	// if len(response.Reviews) == 1 {
	// 	fmt.Println(response)
	// 	t.Errorf("reviews hasn't the correct size")
	// }

	// if !(response.Reviews[0].Resume == "Really awesome" &&
	// 	response.Reviews[0].Rating == 4.7) {
	// 	fmt.Println(response)
	// 	t.Errorf("reviews don't correspond with was expected")
	// }
}

func ReviewDelete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewReviewServiceClient(conn)

	response, err := c.Delete(context.Background(), &proto.Id{Id: reviews.Reviews[0].Id})
	if err != nil {
		t.Errorf("Error to connect and get data")
	}

	if response.Code != 0 {
		t.Errorf("not was deleted")
	}
}

func ReviewGetByBookId(t *testing.T) {
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewReviewServiceClient(conn)

	response, err := c.GetByBookId(context.Background(), &proto.Id{Id: book.Id})
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(response.Reviews) != 2 {
		t.Errorf("reviews hasn't the correct size")
	}

	count := 0
	for _, o := range reviews_request {
		for _, r := range response.Reviews {
			if o.Id == r.Id &&
				o.Resume == r.Resume &&
				o.Rating == r.Rating &&
				o.AccountId == r.AccountId &&
				o.BookId == r.BookId {

				count += 1
				break
			}
		}
	}

	if len(response.Reviews) != 2 {
		fmt.Println(response)
		t.Errorf("reviews don't correspond with was expected")
	}

	reviews = *response
}

func ReviewDeleteTotal(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":9000"), grpc.WithInsecure())

	if err != nil {
		t.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewReviewServiceClient(conn)

	for _, review := range reviews.Reviews {
		response, err := c.Delete(context.Background(), &proto.Id{Id: review.Id})
		if err != nil {
			t.Errorf("Error to connect and get data")
		}

		if response.Code != 0 {
			t.Errorf("not was deleted")
		}
	}
}
