package database

import "github.com/nest-hackathon/server/internal/types"

type Database interface {
	GetUserByUsername(username string) (*types.User, error)
	CreateUser(user *types.User) error
}

type DatabaseImpl struct {
	// Add your database implementation details here
}

func NewDatabase() *DatabaseImpl {
	return &DatabaseImpl{}
}

func (db *DatabaseImpl) GetUserByUsername(username string) (*types.User, error) {
	// Implement database query
	return nil, nil
}

func (db *DatabaseImpl) CreateUser(user *types.User) error {
	// Implement database insert
	return nil
}
