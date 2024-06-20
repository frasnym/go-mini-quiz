package main

/*
1. Iterate the numbers for 1st number
2. Iterate for the 2nd numbers
3. Compare target with sum the number; 1st + 2nd
*/
func twoSum(numbers []int, target int) []int {
	// Iterate the numbers for first number
	for i1, n1 := range numbers {

		// Iterate for the second numbers
		// Make sure to not process same index, therefore we add +1 to the starting index
		for i := i1 + 1; i < len(numbers); i++ {

			// Compare target with sum the number; 1st + 2nd
			if n1+numbers[i] == target {
				return []int{i1, i}
			}

		}
	}
	return []int{}
}
