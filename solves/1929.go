package main

func getConcatenation(nums []int) []int {
    result := make([]int, len(nums)*2)
	index := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < len(nums); j ++ {
			result[index] = nums[j]
			index++
		}
	}
	return result
}