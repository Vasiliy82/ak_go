package main

import "fmt"

func getFibon() func() int {
	var a = 0
	var b = 1
	return func() int {
		result := a
		sum := a + b
		a = b
		b = sum
		return result
	}
}

func main() {
	nextFibon := getFibon()
	arr := [...]int{nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(),
		nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(), nextFibon(),
		nextFibon(), nextFibon(), nextFibon(),
	}

	fmt.Print(arr)

}
