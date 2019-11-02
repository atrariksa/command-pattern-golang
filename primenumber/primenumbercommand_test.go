package primenumber

import (
	"testing"

	"github.com/atrariksa/command-pattern-golang/input"
	"github.com/stretchr/testify/assert"
)

func TestPrimeNumberCommand(t *testing.T) {
	primeNumberInput := input.PrimeNumberInput{8}
	actualPrimeNumberResult := Execute(primeNumberInput)
	expectedResult := []int{2, 3, 5, 7, 11, 13, 17, 19}
	assert.Equal(t, actualPrimeNumberResult.Value, expectedResult, "they should be equal")
}
