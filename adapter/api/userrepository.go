package api

import (
	"github.com/PTH-IT/batchi_golang/domain/repository"
)

func NewUser() repository.UserRepository {
	return userRepository{}
}

type userRepository struct {
}

func (repo userRepository) GetUser() (string, error) {

	return "user", nil
}
