package models

import "method-chaining/database"

type user struct {
	database.Model

	Name     string
	Email    string
	Password string
}

func User(name, email, password string) *user {
	u := new(user)

	u.Model = database.NewModel()
	u.SetTable("users")

	u.Name = name
	u.Email = email
	u.Password = password

	return u
}
