package main

import "fmt"

func FirstReverse(str string) string {

	// code goes here
	result := ""
	for i := len(str); i > 0; i-- {
		result = result + string((str[i-1]))
	}
	return result

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FirstReverse("Hello World and Coders"))

}
