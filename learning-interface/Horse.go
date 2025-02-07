package main

import "fmt"

type Horse struct {
	Name  string
	Color string
	Age   int
}

func (horse Horse) Run() {
	fmt.Println(horse.Name, "is running")
}
