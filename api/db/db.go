package db

import (
	"vote-app/models"
)

type Persist interface {
	Connect() (Persist, error)
	SaveUser(model.User) error
	GetUser() ([]model.User, error)
}
