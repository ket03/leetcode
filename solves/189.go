package main

import "fmt"
// разбить на смещаемое и не смещаемое
func rotate(nums []int, k int)  {
	temp := 0
	reserve := 0
    for i := 0; i < k; i++ {
		temp = nums[i]
		nums[i] = nums[len(nums)-k+i]
		for j := 0; j < len(nums)-k; j++ {
			reserve = nums[j]
			nums[j] = temp
			temp = reserve
			fmt.Println(nums)
		}
		// nums[len(nums)-k+i] = temp
	}
	// fmt.Println(nums)
}