package db

import "github.com/kconde2/vote-app/api/models"

// Local ss
type Local struct{}

var (
	listUser      []User
	incrementUser int
)

// Connect c
func (l Local) Connect() (Persist, error) {
	return nil, nil
}

// User u
type User = models.User

// SaveUser S
func (l Local) SaveUser(u User) error {
	incrementUser++
	u.Id = incrementUser
	listUser = append(listUser, u)
	return nil
}

// GetUser g
func (l Local) GetUser() ([]User, error) {
	return listUser, nil
}
