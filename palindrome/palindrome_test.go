package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("ada"))
	assert.True(t, IsPalindrome(""))
	assert.True(t, IsPalindrome("kodok"))
	assert.True(t, IsPalindrome("madam"))
	assert.True(t, IsPalindrome("racecar"))
	assert.True(t, IsPalindrome("saippuakivikauppias"))

	assert.False(t, IsPalindrome("me"))
	assert.False(t, IsPalindrome("you"))
	assert.False(t, IsPalindrome("you"))
}

func TestIsPalindromeNoNewVar(t *testing.T) {
	assert.True(t, IsPalindromeNoNewVar("ada"))
	assert.True(t, IsPalindromeNoNewVar(""))
	assert.True(t, IsPalindromeNoNewVar("kodok"))
	assert.True(t, IsPalindromeNoNewVar("madam"))
	assert.True(t, IsPalindromeNoNewVar("racecar"))
	assert.True(t, IsPalindromeNoNewVar("saippuakivikauppias"))

	assert.False(t, IsPalindromeNoNewVar("me"))
	assert.False(t, IsPalindromeNoNewVar("you"))
	assert.False(t, IsPalindromeNoNewVar("you"))
}

func TestIsPalindromeNoNewVarMoreOptimal(t *testing.T) {
	assert.True(t, IsPalindromeNoNewVarMoreOptimal("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	assert.True(t, IsPalindromeNoNewVarMoreOptimal("ada"))
	assert.True(t, IsPalindromeNoNewVarMoreOptimal(""))
	assert.True(t, IsPalindromeNoNewVarMoreOptimal("kodok"))
	assert.True(t, IsPalindromeNoNewVarMoreOptimal("madam"))
	assert.True(t, IsPalindromeNoNewVarMoreOptimal("racecar"))
	assert.True(t, IsPalindromeNoNewVarMoreOptimal("saippuakivikauppias"))

	assert.False(t, IsPalindromeNoNewVarMoreOptimal("me"))
	assert.False(t, IsPalindromeNoNewVarMoreOptimal("you"))
	assert.False(t, IsPalindromeNoNewVarMoreOptimal("you"))
}

func TestIsPalindromeRecursive(t *testing.T) {
	assert.True(t, IsPalindromeRecursive("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	assert.True(t, IsPalindromeRecursive("ada"))
	assert.True(t, IsPalindromeRecursive(""))
	assert.True(t, IsPalindromeRecursive("kodok"))
	assert.True(t, IsPalindromeRecursive("madam"))
	assert.True(t, IsPalindromeRecursive("racecar"))
	assert.True(t, IsPalindromeRecursive("saippuakivikauppias"))

	assert.False(t, IsPalindromeRecursive("me"))
	assert.False(t, IsPalindromeRecursive("you"))
	assert.False(t, IsPalindromeRecursive("you"))
}
