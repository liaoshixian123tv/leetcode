package question6

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestConvertA(t *testing.T) {
	result := Convert("pwwkew", 1)
	assert.Equal(t, "pwwkew", result)
}

func TestConvertB(t *testing.T) {
	result := Convert("PAYPALISHIRING", 3)
	assert.Equal(t, "PAHNAPLSIIGYIR", result)
}

func TestConvertC(t *testing.T) {
	result := Convert("PAYPALISHIRING", 4)
	assert.Equal(t, "PINALSIGYAHRPI", result)
}

func TestConvertD(t *testing.T) {
	result := Convert("p", 1)
	assert.Equal(t, "p", result)
}

func TestConvertE(t *testing.T) {
	result := Convert("p", 2)
	assert.Equal(t, "p", result)
}
