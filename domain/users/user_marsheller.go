package users

import "encoding/json"

type PublicUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type PrivateUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	bt, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(bt, &publicUser)
		return publicUser
	}
	var privateUser PrivateUser
	json.Unmarshal(bt, &privateUser)
	return privateUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, v := range users {
		result[i] = v.Marshall(isPublic)
	}
	return result
}
