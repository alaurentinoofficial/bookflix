package db

import (
	"fmt"
	"log"
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

	// db.DropTableIfExists(&models.Catalog{}, &models.Banner{})
	// db.AutoMigrate(&models.Catalog{}, &models.Banner{})
}

func Get() *gorm.DB {
	return db
}
