package sum

import (
	"testing"

	"github.com/atrariksa/command-pattern-golang/input"
	"github.com/stretchr/testify/assert"
)

func TestSumCommand(t *testing.T) {
	sumInput := input.SumInput{8, 11}
	actualSumResult := Execute(sumInput)
	expectedResult := 19
	assert.Equal(t, actualSumResult.Value, expectedResult, "they should be equal")
}
