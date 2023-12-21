package main

func removeDuplicates(nums []int) int {
	current := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[current] = nums[i]
			current++
		}
	}
	return current
}