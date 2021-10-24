package users

import (
	"fmt"

	"github.com/SriHemnath/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	if userDB[user.ID] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
