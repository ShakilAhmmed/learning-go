package models

import "method-chaining/database"

type person struct {
	database.Model
	Name  string
	Email string
}

func Person(name, email string) *person {
	p := new(person)
	p.Model = database.NewModel()
	p.SetTable("persons")

	p.Name = name
	p.Email = email

	return p
}
