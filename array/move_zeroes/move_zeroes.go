package main

/*
1. Count notZero
2. Assign notZero to beginning of array
3. Change remaining array value to 0
*/
func moveZeroes(nums []int) []int {
	notZeroCount := 0
	for _, v := range nums {
		if v != 0 {
			nums[notZeroCount] = v // Assign notZero to beginning of array
			notZeroCount++         // Count notZero
		}
	}

	// Change remaining array value to 0
	for i := notZeroCount; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}

/*
1. Count zero
2. Count notZero
3. Assign notZero to beginning of array
4. Combine notZero array with 0 array
*/
func moveZeroes1(nums []int) []int {
	zeroCount := 0
	notZeroCount := 0
	for _, v := range nums {
		if v == 0 {
			zeroCount++ // Count zero
		} else {
			nums[notZeroCount] = v // Assign notZero to beginning of array
			notZeroCount++         // Count notZero
		}
	}

	// Combine notZero array with 0 array
	nums = append(nums[:notZeroCount], make([]int, zeroCount)...)
	return nums
}
