package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyProcess(t *testing.T) {
	multiplyInput := MultiplyInput{8, 11}
	multiplyResult := Multiply(multiplyInput)
	expectedResult := 88
	assert.Equal(t, multiplyResult, expectedResult, "they should be equal")
}
