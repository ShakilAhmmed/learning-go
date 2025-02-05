package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sum[T Number](numbers []T) (sum T) {
	for _, val := range numbers {
		sum += val
	}
	return
}

func sumInt(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}

func sumInt8(numbers []int8) (sum int8) {
	for _, val := range numbers {
		sum += val
	}
	return
}

func sumFloat32(numbers []float32) (sum float32) {
	for _, val := range numbers {
		sum += val
	}
	return
}
func main() {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	numbers2 := []int8{1, 2, 3, 4, 5, 6}
	numbers3 := []float32{1.1, 1.2, 1.3}

	fmt.Println(sumInt(numbers1))
	fmt.Println(sumInt8(numbers2))
	fmt.Println(sumFloat32(numbers3))

	// refactor with generics
	fmt.Println(sum(numbers1))
	fmt.Println(sum(numbers2))
	fmt.Println(sum(numbers3))

}
