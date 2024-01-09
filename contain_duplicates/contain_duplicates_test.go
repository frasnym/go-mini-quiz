package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tcs := []struct {
		Name        string
		Input       []int
		Expectation bool
	}{
		{"duplicate", []int{1, 2, 3, 1}, true},
		{"no duplicate", []int{1, 2, 3, 4}, false},
		{"duplicate", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range tcs {
		output := containsDuplicate(tc.Input)
		assert.Equal(t, tc.Expectation, output)
	}
}
