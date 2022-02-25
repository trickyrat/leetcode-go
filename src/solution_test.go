package main

import (
	"testing"
)

func TestComplexNumberMultiply(t *testing.T) {
	var (
		num1     = "1+1i"
		num2     = "1+1i"
		expected = "0+2i"
	)
	actual := complexNumberMultiply(num1, num2)
	if actual != expected {
		t.Errorf("ComplexNumberMultiply(%s, %s) = %s; expected %s", num1, num2, actual, expected)
	}
}
