package mymath_test

import (
	"demounit/mymath"
	"testing"
)

// 3A principle
func TestOnePlusTwoEqualsThree(t *testing.T) {
	// Arrange
	num1 := 1
	num2 := 2

	expect := 3 // Expected output of Add(1,2)
	actual := 0 // Actual output

	// Act
	actual = mymath.Add(num1, num2)

	// Assert
	if expect != actual {
		t.Error("1 + 2 should be 3")
	}
}
