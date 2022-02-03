package main

import (
	"fmt"
)

func main() {
	var delete, append int

	delete, append, _ = appendDelete("klikakademi", "klikdokter")
	fmt.Printf("Final Result = Delete: %v, append: %v\n", delete, append)

	delete, append, _ = appendDelete("abc", "def")
	fmt.Printf("Final Result = Delete: %v, append: %v\n", delete, append)

	delete, append, _ = appendDelete2("klikakademi", "klikdokter")
	fmt.Printf("Final Result = Delete: %v, append: %v\n", delete, append)

	delete, append, _ = appendDelete2("abc", "def")
	fmt.Printf("Final Result = Delete: %v, append: %v\n", delete, append)
}

func appendDelete(str1, str2 string) (int, int, error) {
	delete := 0
	append := 0
	newStr1 := ""

	for newStr1 != str2 {
		for i := len(newStr1); i < len(str2); i++ {
			char1 := string([]rune(str1)[i])
			char2 := string([]rune(str2)[i])
			// fmt.Printf("Comparing %v with %v\n", (str1[:i] + "[" + char1 + "]" + str1[i+1:]), (str2[:i] + "[" + char2 + "]" + str2[i+1:]))

			if char1 == char2 {
				newStr1 += char1
			} else {
				str1 = newStr1 + str1[i+1:]
				delete++

				newStr1 = newStr1 + char2
				str1 = str1 + char2
				append++
			}
		}
	}

	return delete, append, nil
}

func appendDelete2(str1, str2 string) (int, int, error) {
	i := 0
	char1 := string([]rune(str1)[i])
	char2 := string([]rune(str2)[i])

	for char1 == char2 {
		i++
		char1 = string([]rune(str1)[i])
		char2 = string([]rune(str2)[i])
	}

	newStr1 := str1[i:]
	newStr2 := str2[i:]

	return len(newStr1), len(newStr2), nil
}
