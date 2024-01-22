package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Name        string
	Input1      []int
	Input2      []int
	Expectation []int
}{
	{"return 2,2", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
	{"return 4,9", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	{"return 1", []int{1, 2}, []int{1, 1}, []int{1}},
	{"return 1 (2)", []int{3, 1, 2}, []int{1, 1}, []int{1}},
}

func TestIntersect(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			actual := intersect(tc.Input1, tc.Input2)
			assert.Equal(t, tc.Expectation, actual)
		})
	}
}
