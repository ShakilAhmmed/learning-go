package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name       string `json:"name"`
	Credit     int    `json:"credit"`
	CreditHour int    `json:"creditHour"`
}
type Student struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Courses []Course `json:"courses"`
}

func main() {
	students := []Student{
		{
			Name:  "Shakil Ahmmed",
			Age:   22,
			Email: "shakil.ahmmed@vub.be",
			Courses: []Course{
				{
					Name:       "Advanced Programming Concept",
					Credit:     3,
					CreditHour: 24,
				},
				{
					Name:       "Biomedical Signal and Images",
					Credit:     3,
					CreditHour: 16,
				},
			},
		},
		{
			Name:  "Tamim Ahmmed",
			Age:   23,
			Email: "tamim.ahmmed@vub.be",
			Courses: []Course{
				{
					Name:       "Declarative Programming",
					Credit:     6,
					CreditHour: 36,
				},
				{
					Name:       "Biomedical Signal and Images",
					Credit:     3,
					CreditHour: 16,
				},
			},
		},
		{
			Name:    "Niloy Deb",
			Age:     23,
			Email:   "niloy.deb@vub.be",
			Courses: []Course{},
		},
	}

	responseJson, error := json.MarshalIndent(students, "", "\t")

	if error != nil {
		panic(error)
	}

	fmt.Printf("%s\n", responseJson)

}
