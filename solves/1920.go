package main

func buildArray(nums []int) []int {
	arr_reserve := make([]int, len(nums))
	copy(arr_reserve, nums)
	for i := range nums {
		nums[i] = arr_reserve[arr_reserve[i]]
	}
	return nums
}