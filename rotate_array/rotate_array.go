package main

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Target: O(1) extra space

*/

// Complexity = O(n)
func rotate2(nums []int, k int) {
	// 0 + 3 = 3 % 7 = 3
	// 1 + 3 = 4 % 7 = 4
	// 2 + 3 = 5 % 7 = 5
	// 3 + 3 = 6 % 7 = 6
	// 4 + 3 = 7 % 7 = 0
	// 5 + 3 = 8 % 7 = 1
	// 6 + 3 = 9 % 7 = 2

	output := make([]int, len(nums))
	for i, v := range nums {
		oIndex := (i + k) % len(nums)
		output[oIndex] = v
	}
	copy(nums, output)
}

// Complexity = O(n)
// Slicing array takes O(n) time since it creates two new arrays based on the original input.
func rotate(nums []int, k int) {
	// [1,2,3,4,5,6,7]
	// [5,6,7,1,2,3,4]

	// [-,-,-,1,2,3,4]
	// [5,6,7,-,-,-,-]

	lenInput := len(nums)

	if k > lenInput {
		k = k % lenInput
	}
	divider := lenInput - k - 1

	end := nums[divider+1:]
	begin := nums[:divider+1]

	output := append(end, begin...)

	copy(nums, output)
}
