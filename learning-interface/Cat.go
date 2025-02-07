package main

import "fmt"

type Cat struct {
	Name  string
	Color string
}

func (cat Cat) Sound() {
	fmt.Println(cat.Name, " with color ", cat.Color, "Sounds like EEEE")
}

func (cat Cat) Move() {
	fmt.Println(cat.Name, " with color ", cat.Color, "Moving....")
}
