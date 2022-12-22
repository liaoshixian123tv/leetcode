package question7

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestReverseA(t *testing.T) {
	result := Reverse(15486106)
	assert.Equal(t, 60168451, result)
}

func TestReverseB(t *testing.T) {
	result := Reverse(-14978416)
	assert.Equal(t, -61487941, result)
}

func TestReverseC(t *testing.T) {
	result := Reverse(456415)
	assert.Equal(t, 514654, result)
}

func TestReverseD(t *testing.T) {
	result := Reverse(1534236469)
	assert.Equal(t, 0, result)
}

func TestReverseE(t *testing.T) {
	result := Reverse(1534236469)
	assert.Equal(t, 0, result)
}
