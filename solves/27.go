package main

import "slices"

func removeElement(nums []int, val int) int {
	counter := 0
	for i := range nums {
		if nums[i] == val {
			nums[i] = 101
		} else {
            counter++
        }
	}
	slices.Sort(nums)
	return counter
}