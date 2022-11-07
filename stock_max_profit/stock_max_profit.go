/**
Array Challenge
Have the function ArrayChallenge(arr) take the array of numbers stored in arr which will contain integers that represent the amount in dollars that a single stock is worth, and return the maximum profit that could have been made by buying stock on day x and selling stock on day y where y > x. For example: if arr is [44, 30, 24, 32, 35, 30, 40, 38, 15] then your program should return 16 because at index 2 the stock was worth $24 and at index 6 the stock was then worth $40, so if you bought the stock at 24 and sold it at 40, you would have made a profit of $16, which is the maximum profit that could have been made with this list of stock prices.

If there is not profit that could have been made with the stock prices, then your program should return -1. For example: arr is [10, 9, 8, 2] then your program should return -1.
Examples
Input: []int {10,12,4,5,9}
Output: 5
Input: []int {14,20,4,12,5,11}
Output: 8
*/

package main

import (
	"fmt"
)

// Solution1 have O(n^2) time complexity
func Solution1(arr []int) int {
	maxProfit := -1
	for index, buy := range arr {

		profit := -1
		for _, sell := range arr[index:] {
			// If buy price is bigger or same, we cannot get profit
			// Proceed to next sell price
			if buy >= sell {
				continue
			}

			profit = sell - buy

			// If the profit is bigger than previous profit
			// Replace the maximum profit
			if profit > maxProfit {
				maxProfit = profit
			}
		}

	}

	return maxProfit
}

// Solution2 have O(n) time complexity
func Solution2(arr []int) int {
	maxProfit := -1
	buyPrice := 0
	sellPrice := 0

	isChangeBuyIndex := true
	for i, v := range arr {

		// If next day if not available, stop loop
		if len(arr) <= i+1 {
			break
		}

		// Sell on the next day after buy
		sellPrice = arr[i+1]

		// Set initial buy price to today
		if isChangeBuyIndex {
			buyPrice = v
		}

		// If sell price is lower, cannot make profit
		// Change the buy index
		if sellPrice < buyPrice {
			isChangeBuyIndex = true
			continue
		} else {
			// If we can make profit
			// Check if profit is bigger that previous profit
			profit := sellPrice - buyPrice
			if profit > maxProfit {
				maxProfit = profit
			}
			isChangeBuyIndex = false
		}
	}

	return maxProfit
}

func main() {
	stockPrice := []int{14, 20, 4, 12, 5, 11}

	fmt.Println(Solution1(stockPrice))
	fmt.Println(Solution2(stockPrice))
	// sell 20
	// buy 14
	// profit 6
	// max profit 6

	// sell 4
	// buy 14

	// sell 12
	// buy 4
	// profit 8
	// max profit 8

	// sell 5
	// buy 4
	// profit 1

	// sell 11
	// buy 4
	// profit 7
}
