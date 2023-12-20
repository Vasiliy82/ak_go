package main

import (
	"fmt"
	"log"
)

type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}
func main() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}
func Report(x any) {
	switch x.(type) {
	case Xml:
		x.(Xml).Format()
	case Csv:
		x.(Csv).Format()
	default:
		log.Fatalln("Неподдерживаемый формат")
	}
}
