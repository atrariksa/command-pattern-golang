package input

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeNumberProcess(t *testing.T) {
	primeNumberInput := PrimeNumberInput{8}
	primeNumberResult, err := PrimeNumber(primeNumberInput)
	if err != nil {
		log.Fatal(err)
		t.Error()
	}
	expectedResult := []int{2, 3, 5, 7, 11, 13, 17, 19}
	assert.Equal(t, primeNumberResult, expectedResult, "they should be equal")
}
