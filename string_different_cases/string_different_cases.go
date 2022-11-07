/**
Different Cases
HIDE QUESTION
Have the function DifferentCases(str) take the str parameter being passed and return it in upper camel case format where the first letter of each word is capitalized. The string will only contain letters and some combination of delimiter punctuation characters separating each word.

For example: if str is "Daniel LikeS-coding" then your program should return the string DanielLikesCoding.
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Solution1 have O(n)
func Solution1(str string) string {

	// Replace all symbol with space using regex
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		fmt.Println(err.Error())
	}
	str = re.ReplaceAllString(str, " ")

	// Convert sentence to lower case
	str = strings.ToLower(str)

	// Transform sentence to array of word
	words := strings.Split(str, " ")

	// Loop for each word
	result := ""
	for _, v := range words {
		// Add current result with word that already Capitalized on first letter
		result = result + strings.Title(v)
	}

	return result
}

func main() {
	fmt.Println(Solution1("Daniel LikeS-coding"))
}
