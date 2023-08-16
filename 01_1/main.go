package main

import "fmt"

func getInc() func() int {
	var a = 0
	return func() int {
		a++
		return a
	}
}

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
	nextCtr := getInc()

	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())
	fmt.Printf("fibon(%d): %d\n", nextCtr(), nextFibon())

}
