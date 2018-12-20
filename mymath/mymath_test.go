package mymath_test

import (
	"demounit/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1 + 2 = 3
func TestOnePlusTwoEqualsThree(t *testing.T) {
	// Arrange
	num1 := 1
	num2 := 2

	expect := 3 // Expected output of Add(1,2)
	actual := 0 // Actual output

	// Act
	actual = mymath.Add(num1, num2)

	// Assert
	assert.Equal(t, expect, actual)
}

// 1 + 1 = 2
func TestOnePlusOneEqualsTwo(t *testing.T) {
	// Arrange
	num1 := 1
	num2 := 1

	expect := 2 // Expected output of Add(1,2)
	actual := 0 // Actual output

	// Act
	actual = mymath.Add(num1, num2)

	// Assert
	assert.Equal(t, expect, actual)
}

// 2 * 3 = 6
func Test2x3Equals6(t *testing.T) {
	// Arrange
	num1 := 2
	num2 := 3

	expect := 6 // Expected output of Add(1,2)
	actual := 0 // Actual output

	// Act
	actual = mymath.Multiply(num1, num2)

	// Assert
	assert.Equal(t, expect, actual)
}

// 3 * 4 = 12
func Test3x4Equals12(t *testing.T) {
	// Arrange
	num1 := 3
	num2 := 4

	expect := 12 // Expected output of Add(1,2)
	actual := 0  // Actual output

	// Act
	actual = mymath.Multiply(num1, num2)

	// Assert
	assert.Equal(t, expect, actual)
}
