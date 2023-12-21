package main

import "fmt"

const (
	n = 10
)

func main() {

	lastResult1 := 1
	lastResult2 := 0
	result := 0
	num := 1

	var nextFibonacci func()

	nextFibonacci = func() {

		_ = (num > 1 && num <= n && func() bool {
			lastResult2 = lastResult1
			lastResult1 = result
			result = lastResult2 + lastResult1
			return true
		}())

		fmt.Printf("Fibonacci[%d] = %d\n", num, result)
		num++
		_ = num <= n && func() bool {
			nextFibonacci()
			return true
		}()
	}

	nextFibonacci()

}
