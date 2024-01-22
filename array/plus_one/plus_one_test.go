package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Name        string
	Input       []int
	Expectation []int
}{
	{"1", []int{1, 2, 3}, []int{1, 2, 4}},
	{"2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{"3", []int{9}, []int{1, 0}},
	{"4", []int{9, 9}, []int{1, 0, 0}},
	{"5", []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}, []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
	{"6", []int{8, 9, 9, 9}, []int{9, 0, 0, 0}},
}

func TestSingleNumber(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			actual := plusOne(tc.Input)
			assert.Equal(t, tc.Expectation, actual)
		})
	}
}
