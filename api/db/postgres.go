package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	POSTGRES_USER     = os.Getenv("PORT")
	POSTGRES_DB       = os.Getenv("POSTGRES_DB")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)

type PostGresSQL struct {
	db *gorm.DB
}

func (p PostGresSQL) Connect() (Persist, error) {
	login := "host=db port=5432 user=" + POSTGRES_USER + " dbname=" + POSTGRES_DB + " password=" + POSTGRES_PASSWORD
	db, err := gorm.Open("postgres", login)
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	db.AutoMigrate(&User{}, &Proposal{})
	p.db = db
	return &p, nil
}

func (p PostGresSQL) SaveUser(u User) error {
	return p.db.Create(&u).Error
}

func (p PostGresSQL) GetUser() ([]User, error) {
	var us []User
	err := p.db.Find(&us).Error
	return us, err
}
