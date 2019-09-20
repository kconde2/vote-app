package db

// Import
import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// Postgres to database connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kconde2/vote-app/api/models"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Initialize creates a connection to postgres database and migrates any new models
func Initialize() {
	user := getEnv("POSTGRES_USER", "userp")
	password := getEnv("POSTGRES_PASSWORD", "password")
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	database := getEnv("POSTGRES_DB", "dev")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	db.LogMode(true)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	log.Println("Database connected")

	// User Table
	if !db.HasTable(&models.User{}) {
		err := db.CreateTable(&models.User{})
		if err != nil {
			log.Println("users ==> Table already exists")
		} else {
			log.Println("users ==> Successfully created")
		}
	}

	// Vote Table
	if !db.HasTable(&models.Vote{}) {
		err := db.CreateTable(&models.Vote{})
		if err != nil {
			log.Println("votes ==> Table already exists")
		} else {
			log.Println("users ==> Successfully created")
		}
	}

	// Run migrations (models field changes)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Vote{})
}

// GetDB gDB
func GetDB() *gorm.DB {
	return db
}

// CloseDB cDB
func CloseDB() {
	db.Close()
}
