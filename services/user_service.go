package services

import (
	"github.com/SriHemnath/bookstore_users-api/domain/users"
	"github.com/SriHemnath/bookstore_users-api/utils/crypto_utils"
	"github.com/SriHemnath/bookstore_users-api/utils/date_utils"
	"github.com/SriHemnath/bookstore_users-api/utils/errors"
)

var (
	UserService userService = &us{}
)

type us struct{}

type userService interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(users.User) *errors.RestError
	FindByStatus(string) (users.Users, *errors.RestError)
}

func (s *us) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *us) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *us) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
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

func (s *us) DeleteUser(user users.User) *errors.RestError {
	//result := &users.User{ID: userId}
	if err := user.GetUserByEmail(); err != nil {
		return err
	}
	return user.Delete()
}

func (s *us) FindByStatus(status string) (users.Users, *errors.RestError) {
	dto := &users.User{}
	return dto.GetByStatus(status)
}
