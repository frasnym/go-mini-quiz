package main

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1


Constraints:

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

func singleNumber(nums []int) int {
	mapNum := map[int]int{}
	for _, v := range nums {
		mapNum[v] += 1
	}

	for k, v := range mapNum {
		if v == 1 {
			return k
		}
	}

	return 0
}

/*
Bitwise XOR (exclusive or) is a binary operation that takes two equal-length binary representations and performs the XOR operation on each pair of corresponding bits. The result is a new binary number where each bit is the result of the XOR operation between the corresponding bits of the two input numbers.

The XOR operation is denoted by the symbol ^. Here's how it works with individual bits:

0 XOR 0 results in 0.
1 XOR 0 results in 1.
0 XOR 1 results in 1.
1 XOR 1 results in 0.
For example, if you have the binary numbers 1010 and 1100 and you perform a bitwise XOR operation:

markdown
Copy code

	1010

XOR 1100
-------

	0110

So, the result of the bitwise XOR operation in this case is 0110 in binary, which is 6 in decimal.
*/
func singleNumberXor(nums []int) int {
	result := 0

	for _, v := range nums {
		result = result ^ v
	}

	return result
}
