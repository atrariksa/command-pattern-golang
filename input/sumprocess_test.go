package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumProcess(t *testing.T) {
	sumInput := SumInput{8, 11}
	sumResult := Sum(sumInput)
	expectedResult := 19
	assert.Equal(t, sumResult, expectedResult, "they should be equal")
}
