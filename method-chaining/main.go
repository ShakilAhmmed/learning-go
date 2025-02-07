package main

import (
	"method-chaining/database"
	"method-chaining/models"
)

func main() {

	Database := database.MySQL()

	person := models.Person("Shakil Ahmmed", "shakil@gmail.com", Database)
	user := models.User("Shakil Ahmmed", "shakil@gmail.com", "shakil@gmail.com", Database)

	person.Query().Where("name", "shakil").Get()
	user.Query().Where("email", "shakil@gmail.com").Get()
}
