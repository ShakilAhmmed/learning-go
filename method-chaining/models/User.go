package models

import "method-chaining/database"

type user struct {
	database.DatabaseContract

	Name     string
	Email    string
	Password string
}

func User(name, email, password string, db database.DatabaseContract) *user {
	u := new(user)
	u.Name = name
	u.Email = email
	u.Password = password
	u.DatabaseContract = db
	u.SetTable("users")

	return u
}
