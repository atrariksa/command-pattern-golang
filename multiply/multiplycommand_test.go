package multiply

import (
	"testing"

	"github.com/atrariksa/command-pattern-golang/input"
	"github.com/stretchr/testify/assert"
)

func TestMultiplyCommand(t *testing.T) {
	multiplyInput := input.MultiplyInput{8, 11}
	actualMultiplyResult := Execute(multiplyInput)
	expectedResult := 88
	assert.Equal(t, actualMultiplyResult.Value, expectedResult, "they should be equal")
}
