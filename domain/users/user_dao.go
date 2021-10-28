package users

import (
	"fmt"
	"log"
	"strings"

	"github.com/SriHemnath/bookstore_users-api/datasource/mysql/user_db"
	"github.com/SriHemnath/bookstore_users-api/utils/date_utils"
	"github.com/SriHemnath/bookstore_users-api/utils/errors"
	"github.com/SriHemnath/bookstore_users-api/utils/mysql_utils"
)

var (
	userDB = make(map[int64]*User)
)

const (
	email_UNIQUE    = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryGetByEmail = "SELECT id FROM users WHERE email=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?"
)

func (user *User) Get() *errors.RestError {

	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		log.Println("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		log.Println(getErr)
		return mysql_utils.ParseError(getErr)
	}

	//removing mock db
	// result := userDB[user.ID]
	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	// }
	// user.ID = result.ID
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated

	return nil
}

func (user *User) GetUserByEmail() *errors.RestError {

	stmt, err := user_db.Client.Prepare(queryGetByEmail)
	if err != nil {
		log.Println("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID); getErr != nil {
		log.Println(getErr)
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestError {

	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		log.Println("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), email_UNIQUE) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		log.Println("Error while saving the user ", err)
		return mysql_utils.ParseError(err)
	}
	//stmt.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	user.ID, err = result.LastInsertId()
	if err != nil {
		log.Println("Error while saving the user. ID not generated ", err)
		return errors.NewInternalServerError("error when tying to save user")
	}

	//removed mock db
	// if userDB[user.ID] != nil {
	// 	return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.ID))
	// }
	// userDB[user.ID] = user
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		log.Println("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		log.Println("Error while updating user ", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		log.Println("Error while creating statement", err)
		return errors.NewInternalServerError("Not able to save user")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		log.Println("Error while updating user ", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}
