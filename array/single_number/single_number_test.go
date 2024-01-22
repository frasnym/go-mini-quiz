package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Name        string
	Input       []int
	Expectation int
}{
	{"return 1", []int{2, 2, 1}, 1},
	{"return 4", []int{4, 1, 2, 1, 2}, 4},
	{"return 1", []int{1}, 1},
}

func TestSingleNumber(t *testing.T) {
	for _, tc := range tcs {
		actual := singleNumber(tc.Input)
		assert.Equal(t, tc.Expectation, actual)
	}
}

func TestSingleNumberXOR(t *testing.T) {
	for _, tc := range tcs {
		actual := singleNumberXor(tc.Input)
		assert.Equal(t, tc.Expectation, actual)
	}
}
