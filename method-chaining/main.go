package main

func main() {

	Database := &MySQL{}

	person := Person{
		Name:             "Shakil Ahmmed",
		Email:            "shakilfci461@gmail.com",
		DatabaseContract: Database,
	}

	user := User{
		Name:             "Shakil Ahmmed",
		Email:            "shakilfci461@gmail.com",
		Password:         "shakilfci461@gmail.com",
		DatabaseContract: Database,
	}

	person.Query("persons").Where("name", "shakil").Get()
	user.Query("users").Where("email", "shakil@gmail.com").Get()
}
