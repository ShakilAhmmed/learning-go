package main

type Circus struct {
	Name     string
	Location string
}

func (circus Circus) PlayWithAnimal(animalContract AnimalContract) {
	animalContract.Sound()
	animalContract.Move()
}

func main() {
	dog := Dog{
		Name: "Lora",
	}

	cat := Cat{
		Name:  "Dora",
		Color: "Golden",
	}

	//horse := Horse{
	//	Name:  "Rocket",
	//	Color: "Black",
	//	Age:   30,
	//}

	circus := Circus{
		Name:     "2025 New Year Circus",
		Location: "Mechelen 2800",
	}
	circus.PlayWithAnimal(dog)
	circus.PlayWithAnimal(cat)
	//circus.PlayWithAnimal(horse)
}
