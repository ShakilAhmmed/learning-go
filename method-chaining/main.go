package main

import (
	"method-chaining/database"
	"method-chaining/models"
)

func main() {

	//Database := database.MySQL()
	database.InitDatabase(database.MYSQL)

	person := models.Person("Shakil Ahmmed", "shakil@gmail.com")
	user := models.User("Shakil Ahmmed", "shakil@gmail.com", "shakil@gmail.com")

	person.Query().Where("name", "shakil").Get()
	user.Query().Where("email", "shakil@gmail.com").Get()
}
