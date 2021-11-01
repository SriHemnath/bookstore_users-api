package users

import (
	"fmt"
	"strings"

	"github.com/SriHemnath/bookstore_users-api/datasource/mysql/user_db"
	"github.com/SriHemnath/bookstore_users-api/utils/errors"
	"github.com/SriHemnath/bookstore_users-api/utils/logger"
	"github.com/SriHemnath/bookstore_users-api/utils/mysql_utils"
)

var (
	userDB = make(map[int64]*User)
)

const (
	email_UNIQUE     = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser     = "SELECT first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	queryGetByEmail  = "SELECT id FROM users WHERE email=?;"
	queryGetByStatus = "SELECT first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryDeleteUser  = "DELETE FROM users WHERE id=?"
)

func (user *User) Get() *errors.RestError {

	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to get user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error while getting result ", getErr)
		return mysql_utils.ParseError(getErr)
	}

	/*removing mock db
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated */

	return nil
}

func (user *User) GetUserByEmail() *errors.RestError {

	stmt, err := user_db.Client.Prepare(queryGetByEmail)
	if err != nil {
		logger.Error("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to get user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID); getErr != nil {
		logger.Error("Error while getting result", getErr)
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestError {

	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), email_UNIQUE) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		logger.Error("Error while saving the user ", err)
		return mysql_utils.ParseError(err)
	}
	//stmt.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	user.ID, err = result.LastInsertId()
	if err != nil {
		logger.Error("Error while saving the user. ID not generated ", err)
		return errors.NewInternalServerError("error when tying to save user")
	}

	/*removed mock db
	if userDB[user.ID] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.ID))
	}
	userDB[user.ID] = user
	*/
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID, user.Status)
	if err != nil {
		logger.Error("Error while updating user ", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("Error while deleting user ", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

//get all active user
func (user *User) GetByStatus(status string) ([]User, *errors.RestError) {

	stmt, err := user_db.Client.Prepare(queryGetByStatus)
	if err != nil {
		logger.Error("Error while creating statement", err)
		return nil, errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	results, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error while creating statement", err)
		return nil, errors.NewInternalServerError("Not able to get users")
	}
	defer results.Close()

	users := make([]User, 0)
	for results.Next() {
		var user User
		if err := results.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.NewNotFoundError("no active user found")
	}

	return users, nil
}
