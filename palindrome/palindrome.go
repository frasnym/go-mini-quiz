package main

// time complexity: O(n)
func IsPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr = reversedStr + string(str[i])
	}

	return str == reversedStr
}

// time complexity: O(n)
// better space complexity
func IsPalindromeNoNewVar(str string) bool {
	for i := 0; i <= len(str)-1; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

// time complexity: O(n/2)
// better space complexity
func IsPalindromeNoNewVarMoreOptimal(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		// fmt.Println(i, ":", len(str)-i-1)
		// fmt.Println(str[i], ":", str[len(str)-i-1])
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func IsPalindromeRecursiveHelper(str string, i int) bool {
	if i < len(str)/2 {
		if str[i] != str[len(str)-i-1] {
			return false
		}
		return IsPalindromeRecursiveHelper(str, i+1)
	}
	return true
}

func IsPalindromeRecursive(str string) bool {
	return IsPalindromeRecursiveHelper(str, 0)
}
