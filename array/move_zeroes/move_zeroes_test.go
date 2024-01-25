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
	{"1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{"2", []int{0}, []int{0}},
	{"3", []int{0, 0, 1}, []int{1, 0, 0}},
}

func TestSingleNumber(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			result := moveZeroes(tc.Input)
			assert.Equal(t, tc.Expectation, result)
		})
	}
}
