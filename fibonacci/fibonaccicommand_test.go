package fibonacci

import (
	"testing"

	"github.com/atrariksa/command-pattern-golang/input"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciCommand(t *testing.T) {
	fibonacciRecursiveInput := input.FibonacciInput{8, true}
	fibonacciNonRecursiveInput := input.FibonacciInput{8, false}
	actualRecursiveResult := Execute(fibonacciRecursiveInput)
	actualNonRecursiveResult := Execute(fibonacciNonRecursiveInput)
	expectedResult := []int{0, 1, 1, 2, 3, 5, 8, 13}
	assert.Equal(t, actualRecursiveResult.Value, expectedResult, "they should be equal")
	assert.Equal(t, actualNonRecursiveResult.Value, expectedResult, "they should be equal")
}
