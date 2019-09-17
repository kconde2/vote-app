package db

import "vote-app/handlers"

type Local struct{}

var (
	listUser          []User
	incrementUser     int
	listProposal      []Proposal
	incrementProposal int
)

func (l Local) Connect() (Persist, error) {
	return nil, nil
}

type User = model.User

func (l Local) SaveUser(u User) error {
	incrementUser++
	u.Id = incrementUser
	listUser = append(listUser, u)
	return nil
}
func (l Local) GetUser() ([]User, error) {
	return listUser, nil
}
