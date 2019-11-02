package input

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciProcess(t *testing.T) {
	fibonacciRecursiveInput := FibonacciInput{8, true}
	fibonacciNonRecursiveInput := FibonacciInput{8, false}
	actualRecursiveResult, err1 := Fibonacci(fibonacciRecursiveInput)
	if err1 != nil {
		log.Fatal(err1)
		t.Error()
	}
	actualNonRecursiveResult, err2 := Fibonacci(fibonacciNonRecursiveInput)
	if err2 != nil {
		log.Fatal(err2)
		t.Error()
	}
	expectedResult := []int{0, 1, 1, 2, 3, 5, 8, 13}
	assert.Equal(t, actualRecursiveResult, expectedResult, "they should be equal")
	assert.Equal(t, actualNonRecursiveResult, expectedResult, "they should be equal")
}
