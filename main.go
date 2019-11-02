package main

import (
	"fmt"

	"github.com/atrariksa/command-pattern-golang/fibonacci"
	"github.com/atrariksa/command-pattern-golang/input"
	"github.com/atrariksa/command-pattern-golang/multiply"
	"github.com/atrariksa/command-pattern-golang/primenumber"
	"github.com/atrariksa/command-pattern-golang/sum"
)

func main() {
	sumInput := input.SumInput{9, 8}
	fmt.Println(sum.Execute(sumInput).Value)
	multiplyInput := input.MultiplyInput{9, 8}
	fmt.Println(multiply.Execute(multiplyInput).Value)
	primeNumberInput := input.PrimeNumberInput{7}
	fmt.Println(primenumber.Execute(primeNumberInput).Value)
	fibonacciRecursiveInput := input.FibonacciInput{7, true}
	fmt.Println(fibonacci.Execute(fibonacciRecursiveInput).Value)
	fibonacciNotRecursiveInput := input.FibonacciInput{7, false}
	fmt.Println(fibonacci.Execute(fibonacciNotRecursiveInput).Value)
}
