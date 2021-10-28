package users

import (
	"log"

	"github.com/SriHemnath/bookstore_users-api/datasource/mysql/user_db"
	"github.com/SriHemnath/bookstore_users-api/utils/date_utils"
	"github.com/SriHemnath/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT first_name, last_name, email, date_created FROM users WHERE id=?;"
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
		log.Println("Error getting user", err)
		return errors.NewInternalServerError("User Not found")
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
		log.Println("Error while saving the user ", err)
		return errors.NewInternalServerError("error when tying to save user")
	}

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
