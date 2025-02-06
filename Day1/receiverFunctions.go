package main

import "fmt"

type Address struct {
	street      string
	houseNumber int
}
type User struct {
	name    string
	email   string
	phone   string
	address map[string]Address
}

func (user *User) play() {
	fmt.Println(user.name, "is playing")
}

func (user *User) move(position, direction string) {

	if address, exists := user.address["presentAddress"]; exists {
		address.street = direction
		address.houseNumber = 343433
		user.address["presentAddress"] = address
	}
	fmt.Println(user.name, "is moving from", position, "to", direction)
}

//func main() {
//	userOne := User{
//		name:  "Shakil Ahmmed",
//		phone: "00000022121021",
//		email: "shakilf@gmail.com",
//		address: map[string]Address{
//			"presentAddress": {
//				street:      "BefferStraat",
//				houseNumber: 34,
//			},
//			"permanentAddress": {
//				street:      "Feni",
//				houseNumber: 1,
//			},
//		},
//	}
//
//	userOne.play()
//	userOne.move("A", "B")
//	fmt.Println(userOne)
//	fmt.Println(userOne.name)
//}
