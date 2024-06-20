package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Name        string
	Numbers     []int
	Target      int
	Expectation []int
}{
	{"1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"2", []int{3, 2, 4}, 6, []int{1, 2}},
	{"3", []int{3, 3}, 6, []int{0, 1}},
	// Additional test cases
	{"4", []int{1, 2, 3, 4, 5, 6}, 11, []int{4, 5}},
	{"5", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
	{"6", []int{0, 4, 3, 0}, 0, []int{0, 3}},
	{"7", []int{1, -1, 2, -2, 3, -3}, -5, []int{3, 5}},
	{"8", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	{"9", []int{1, 1, 1, 1, 1, 2, 2}, 3, []int{0, 5}},
	{"10", []int{5, 75, 25}, 100, []int{1, 2}},
	{"11", []int{2, 8, 12, 15}, 20, []int{1, 2}},
	{"12", []int{2, 5, 5, 11}, 10, []int{1, 2}},
	{"13", []int{10, 7, 19, 15}, 26, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			result := twoSum(tc.Numbers, tc.Target)
			assert.Equal(t, tc.Expectation, result)
		})
	}
}
