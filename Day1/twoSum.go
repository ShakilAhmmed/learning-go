package main

func twoSum(nums []int, target int) []int {

	bucket := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		find := target - nums[i]
		if _, ok := bucket[find]; ok {
			return []int{bucket[find], i}
		}
		bucket[nums[i]] = i
	}
	return nil
}
