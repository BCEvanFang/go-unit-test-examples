package mymath_test

import (
	"demounit/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyMathModules(t *testing.T) {

	t.Run("1 + 2 should be 3", func(t *testing.T) {
		expect := 3
		actual := mymath.Add(1, 2)
		assert.Equal(t, expect, actual)
	})

	t.Run("1 + 1 should be 2", func(t *testing.T) {
		expect := 2
		actual := mymath.Add(1, 1)
		assert.Equal(t, expect, actual)
	})

	t.Run("2 x 3 should be 6", func(t *testing.T) {
		expect := 6
		actual := mymath.Multiply(2, 3)
		assert.Equal(t, expect, actual)
	})

	t.Run("3 x 4 should be 12", func(t *testing.T) {
		expect := 12
		actual := mymath.Multiply(3, 4)
		assert.Equal(t, expect, actual)
	})
}
