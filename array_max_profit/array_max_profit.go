package main

import (
	"fmt"
)

func ArrayChallenge(arr []int) int {

	// code goes here

	result := -1
	for index, buy := range arr {

		tempResult := -1
		for _, sell := range arr[index:] {
			if buy >= sell {
				continue
			}
			tempResult = sell - buy
			if tempResult > result {
				result = tempResult
			}
		}

	}

	return result

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(ArrayChallenge([]int{14, 20, 4, 12, 5, 11}))

}
