package main

import "fmt"

const (
	n = 20
)

func main() {
	result := make([]int, n)

	result[0], result[1] = 0, 1

	var calcNext func(int) bool

	calcNext = func(id int) bool {
		result[id] = result[id-1] + result[id-2]
		_ = (id < n-1 && calcNext(id+1))
		return true
	}

	_ = calcNext(2)

	fmt.Print(result)
}
