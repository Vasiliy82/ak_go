package main

import (
	"fmt"
	"runtime"
)

func storeValueInFunc(i []int) func() []int {
	return func() []int {
		return i
	}
}
func stat() {
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)

	fmt.Printf("\tAlloc = %v\n", mem.Alloc)
	fmt.Printf("\tTotalAlloc = %v\n", mem.TotalAlloc)
	fmt.Printf("\tSys = %v\n", mem.Sys)
	fmt.Printf("\tNumGC = %v\n", mem.NumGC)
}

func main() {
	funcs := make([]func() []int, 0, 1024)
	for i := 0; i < 1024; i++ {
		arr := make([]int, 1024, 1024)
		funcs = append(funcs, storeValueInFunc(arr))
	}
	fmt.Println("память захвачена")
	stat()
	fmt.Scanf("\n")
	funcs = nil
	fmt.Println("память освобождена")
	stat()
	fmt.Scanf("\n")
	runtime.GC()
	fmt.Println("после runtime.GC()")
	stat()
	fmt.Scanf("\n")

}
