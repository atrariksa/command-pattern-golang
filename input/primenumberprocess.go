package input

import (
	"errors"
	"github.com/atrariksa/command-pattern-golang/constants"
	errorhandler "github.com/atrariksa/command-pattern-golang/errors"
)

func PrimeNumber(input PrimeNumberInput) ([]int, errorhandler.Err) {
	output := []int{}
	if input.X < 1 {
		return output, errors.New(constants.PrimeNumberInputError)
	}
	index := 2
	for {
		if len(output) == input.X {
			break
		}
		if isPrimeNumber(index) {
			output = append(output, index)
		}
		index = index + 1
	}
	return output, nil
}

func isPrimeNumber(input int) bool {
	if input == 2 {
		return true
	}
	if input%2 == 0 {
		return false
	}
	half := input / 2
	for i := 2; i < half; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
