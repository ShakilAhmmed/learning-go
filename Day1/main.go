package main

import (
	"day-1/collections"
	"fmt"
)

func add(firstNumber int, secondNumber int) (sum int) {
	sum = firstNumber + secondNumber
	return sum
}

type UserInfo struct {
	firstName string
	lastName  string
	email     string
	addresses map[string]UserAddress
}

type UserAddress struct {
	street      string
	houseNumber int
}

func main() {
	//var number int = 100
	//secondNumber := 0
	//fmt.Printf("Day 1: %d First Number\n", number)
	//secondNumber = 200
	//fmt.Printf("Day 1: %d Second Number\n", secondNumber)
	//age := 20
	//
	//if age > 18 {
	//	fmt.Println("You are above 18")
	//} else {
	//	fmt.Println("You are below 18")
	//}
	//
	//switch age {
	//case 18:
	//case 17:
	//	fmt.Println("You are below 18")
	//	break
	//case 20:
	//	fmt.Println("You are equal to 20")
	//	break
	//default:
	//	fmt.Println("You are not equal to 20")
	//}
	//
	//fmt.Println(add(20, 30))
	//
	//for i := 0; i < 10; i++ {
	//	//fmt.Printf("Iteration: %v is %v", i, i)
	//	fmt.Println("Iteration is : ", i)
	//}
	//
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println("Sum is:", sum)
	//
	//names := []string{"John", "Paul", "George", "Ringo"}
	//ages := []int{10, 20, 30, 40, 50}
	//
	//for i := 0; i < len(ages); i++ {
	//	fmt.Println(ages[i])
	//}
	//for _, name := range names {
	//	fmt.Println(name)
	//}
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	//val := 42
	//if val := 10; val > 5 {
	//	fmt.Printf("The value of val is %d\n", val)
	//}
	//fmt.Printf("The value of val is %d\n", val)
	//
	//k := 6
	//switch k {
	//case 4:
	//	fmt.Println("was <= 4")
	//	fallthrough
	//case 5:
	//	fmt.Println("was <= 5")
	//	fallthrough
	//case 6:
	//	fmt.Println("was <= 6")
	//	fallthrough
	//case 7:
	//	fmt.Println("was <= 7")
	//	fallthrough
	//case 8:
	//	fmt.Println("was <= 8")
	//	fallthrough
	//default:
	//	fmt.Println("default case")
	//}

	firstUser := UserInfo{
		firstName: "Shakil",
		lastName:  "Ahmmed",
		email:     "shakilfci461@gmail.com",
		addresses: map[string]UserAddress{
			"presentAddress": {
				street:      "BefferStraat",
				houseNumber: 34,
			},
			"permanentAddress": {
				street:      "ABC",
				houseNumber: 343,
			},
		},
	}

	fmt.Println(firstUser)

	fmt.Println(collections.Map([]int{1, 2, 3, 4}, func(index, value int) int {
		return value * 2
	}))

	fmt.Println(collections.Filter([]int{1, 2, 3, 4, 6, 7, 8}, func(index, value int) bool {
		return value%2 != 0
	}))

	fmt.Println(collections.Reject([]int{1, 2, 3, 4, 5, 6, 7}, func(index, value int) bool {
		return value%2 == 0
	}))
}
