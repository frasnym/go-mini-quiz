package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Name        string
	Input       []int
	Step        int
	Expectation []int
}{
	{"rotate by 3", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{"rotate by 2", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	{"rotate by 0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
	{"rotate by length", []int{1, 2, 3}, 3, []int{1, 2, 3}},
	{"rotate by more than length", []int{1, 2, 3}, 5, []int{2, 3, 1}},
}

func TestRotate(t *testing.T) {
	for _, tc := range testCases {
		t.Run("rotate", func(t *testing.T) {
			t.Run(tc.Name, func(t *testing.T) {
				rotate(tc.Input, tc.Step)
				if !reflect.DeepEqual(tc.Input, tc.Expectation) {
					t.Errorf("rotate(%v, %d) got %v, want %v", tc.Input, tc.Step, tc.Input, tc.Expectation)
				}
			})
		})
	}
}

func TestRotate2(t *testing.T) {
	for _, tc := range testCases {
		t.Run("rotate2", func(t *testing.T) {
			t.Run(tc.Name, func(t *testing.T) {
				rotate2(tc.Input, tc.Step)
				if !reflect.DeepEqual(tc.Input, tc.Expectation) {
					t.Errorf("rotate(%v, %d) got %v, want %v", tc.Input, tc.Step, tc.Input, tc.Expectation)
				}
			})
		})
	}
}
