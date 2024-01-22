package main

/*
1. Loop from last index
2. Increase num by 1
3. Check if 10, carry into the previous number
4. If there is no previous number, append 0 to array
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1 // increase number

		if digits[i] == 10 {
			digits[i] = 0 // carry into the previous number
			if i == 0 {   // if there is no previous number, append 0 to array
				digits = append(digits, 0)
				digits[i] = 1
			}
		} else {
			break
		}
	}

	return digits
}
