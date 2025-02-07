package main

import "fmt"

type Dog struct {
	Name string
}

func (dog Dog) Sound() {
	fmt.Println(dog.Name, "Sounds like ABC")
}

func (dog Dog) Move() {
	fmt.Println(dog.Name, "Moving.....")
}
