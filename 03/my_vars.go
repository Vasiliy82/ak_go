package main

import "fmt"

var myVar int
var v1 int = 10
var v2 float32 = 3.26
var v3 = true
var v4 = "my first\tstring"
var v5 = `my first\tstring`
var v6 = 'л'

func main() {
	// fmt.Println(myVar)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println("v4", v4)
	fmt.Println("v5", v5)
	fmt.Println("v6", v6)

}
