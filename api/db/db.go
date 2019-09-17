package db

// Import
import "github.com/kconde2/vote-app/api/models"

// Persist p
type Persist interface {
	Connect() (Persist, error)
	SaveUser(models.User) error
	GetUser() ([]models.User, error)
}
