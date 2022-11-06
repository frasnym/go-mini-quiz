package main

import (
	"fmt"
	"regexp"
	"strings"
)

func StringChallenge(str string) string {

	// code goes here
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		fmt.Println(err.Error())
	}
	str = re.ReplaceAllString(str, " ")
	str = strings.ToLower(str)

	words := strings.Split(str, " ")
	result := ""
	for _, v := range words {
		result = result + strings.Title(v)
	}

	return result
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(StringChallenge("Daniel LikeS-coding"))

}
