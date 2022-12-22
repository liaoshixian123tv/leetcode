package question9

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestIsPalindromeA(t *testing.T) {
	result := IsPalindrome(15486106)
	assert.Equal(t, false, result)
}

func TestIsPalindromeB(t *testing.T) {
	result := IsPalindrome(-15486106)
	assert.Equal(t, false, result)
}

func TestIsPalindromeC(t *testing.T) {
	result := IsPalindrome(-123456789)
	assert.Equal(t, false, result)
}

func TestIsPalindromeD(t *testing.T) {
	result := IsPalindrome(123454321)
	assert.Equal(t, true, result)
}
