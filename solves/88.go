package main

import "slices"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	current := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[current]
		current++
	}
	slices.Sort(nums1)
}