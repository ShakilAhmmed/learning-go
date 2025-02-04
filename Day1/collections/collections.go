package collections

func Map(nums []int, callback func(index, value int) int) []int {
	result := make([]int, 0)
	for index, value := range nums {
		result = append(result, callback(index, value))
	}
	return result
}

func Filter(nums []int, callback func(index, value int) bool) []int {
	result := make([]int, 0)
	for index, value := range nums {
		if callback(index, value) {
			result = append(result, value)
		}
	}
	return result
}

func Reject(nums []int, callback func(index, value int) bool) []int {
	result := make([]int, 0)
	for index, value := range nums {
		if !callback(index, value) {
			result = append(result, value)
		}
	}
	return result
}
