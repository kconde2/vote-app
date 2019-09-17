package db

import (
	"os"

	"github.com/jinzhu/gorm"
	// import
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Environnement variable
var (
	PostgresUser     = os.Getenv("PORT")
	PostgresDb       = os.Getenv("POSTGRES_DB")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
)

// PostGresSQL p
type PostGresSQL struct {
	db *gorm.DB
}

// Connect c
func (p PostGresSQL) Connect() (Persist, error) {
	login := "host=db port=5432 user=" + PostgresUser + " dbname=" + PostgresDb + " password=" + PostgresPassword
	db, err := gorm.Open("postgres", login)
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	p.db = db
	return &p, nil
}

// SaveUser c
func (p PostGresSQL) SaveUser(u User) error {
	return p.db.Create(&u).Error
}

// GetUser c
func (p PostGresSQL) GetUser() ([]User, error) {
	var us []User
	err := p.db.Find(&us).Error
	return us, err
}
