package services

import (
	"github.com/SriHemnath/bookstore_users-api/domain/users"
	"github.com/SriHemnath/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}

func FindUser() {}

func GetUser() {

}

func DeleteUser() {}
