package collections

func Map(nums []int, doMap func(index, value int) int) []int {
	result := make([]int, 0)
	for index, value := range nums {
		result = append(result, doMap(index, value))
	}
	return result
}

func Filter(nums []int, doFilter func(index, value int) bool) []int {
	result := make([]int, 0)
	for index, value := range nums {
		if doFilter(index, value) {
			result = append(result, value)
		}
	}
	return result
}

func Reject(nums []int, doReject func(index, value int) bool) []int {
	result := make([]int, 0)
	for index, value := range nums {
		if !doReject(index, value) {
			result = append(result, value)
		}
	}
	return result
}
