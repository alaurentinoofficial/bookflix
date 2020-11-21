package db

import (
	"fmt"
	"log"
	"ng-book-service/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Init() {

	e := godotenv.Load()
	if e != nil {
		log.Print(e)
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("host='%s' port='%s' user='%s' dbname='%s' password='%s' sslmode=disable", host, port, user, name, pass)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Println("[*] Postgresql error: ", err)
		os.Exit(1)
	} else {
		log.Println("[*] Postgresql connected!")
	}

	db = conn

	db.DropTableIfExists("book_categories", &models.Category{}, &models.Book{})
	db.AutoMigrate(&models.Category{}, &models.Book{})
	db.Table("book_categories").AddForeignKey("book_id", "books(id)", "CASCADE", "CASCADE")
	db.Table("book_categories").AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
}

func Get() *gorm.DB {
	return db
}
