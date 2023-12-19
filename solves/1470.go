package main

func shuffle(nums []int, n int) []int {
	current := 0
	var result []int
    for i := 0; i < len(nums); i++ {
		if (i % 2 != 0 && i < n) {
			result[i] = nums[n+i-1]
		} else {
			result[i] = nums[current]
			current++
		}
	}
	return nums
}