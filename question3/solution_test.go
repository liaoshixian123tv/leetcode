package question3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongSubstringA(t *testing.T) {
	result := LengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, result)
}

func TestLongSubstringB(t *testing.T) {
	result := LengthOfLongestSubstring("pwwkewadrt")
	assert.Equal(t, 7, result)
}

func TestLongSubstringC(t *testing.T) {
	result := LengthOfLongestSubstring("ppppp")
	assert.Equal(t, 1, result)
}
