package main

import (
	"fmt"
)

func fibon(n int) int {
	switch n {
	case 1:
		return 0
	case 2:
		return 1
	default:
		return fibon(n-2) + fibon(n-1)
	}
}
func seq(n int) int {
	switch n {
	case 24:
		return 24
	default:
		fmt.Printf("fibon(%d) = %d\n", n, fibon(n))
		return seq(n + 1)
	}
}

func main() {
	_ = seq(1)
}
