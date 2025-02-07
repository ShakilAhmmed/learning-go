package models

import "method-chaining/database"

type person struct {
	database.DatabaseContract

	Name  string
	Email string
}

func Person(name, email string, db database.DatabaseContract) *person {
	p := new(person)
	p.Name = name
	p.Email = email
	p.DatabaseContract = db
	p.SetTable("persons")
	return p
}
