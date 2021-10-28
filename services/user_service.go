package services

import (
	"github.com/SriHemnath/bookstore_users-api/domain/users"
	"github.com/SriHemnath/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUser() {}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	currentUser := &users.User{ID: user.ID}
	if err := currentUser.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}

		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}

		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if updateErr := currentUser.Update(); updateErr != nil {
		return nil, updateErr
	}
	return currentUser, nil
}

func DeleteUser(user users.User) *errors.RestError {
	//result := &users.User{ID: userId}
	if err := user.GetUserByEmail(); err != nil {
		return err
	}
	return user.Delete()
}
