package main

func removeDuplicatesII(nums []int) int {
	counter := 0
	isCorrect := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] && isCorrect {
			nums[counter] = nums[i]
			nums[counter+1] = nums[i+1]
			counter += 2
			isCorrect = false
		}
		if nums[i] != nums[i+1] {
			nums[counter] = nums[i]
			isCorrect = true
		}
	}

	return counter
}
