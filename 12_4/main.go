package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	message string
}

func (e myFirstError) Error() string {
	return e.message
}

func main() {
	// кейс, если ошибка была
	err1 := myFirstError{message: "ошибка1"}
	// кейс, если не было
	// err1 := errors.New("ошибка1")

	err2 := fmt.Errorf("ошибка2:%w", err1)
	err3 := fmt.Errorf("ошибка3:%w", err2)

	fmt.Println(err3)

	// Не используя unwrap, нужно определить была ли ошибка «ошибка1»
	if found := errors.As(err3, new(myFirstError)); found == true {
		fmt.Println("Вложенная ошибка \"ошибка1\" была")
	} else {
		fmt.Println("Вложенной ошибки \"ошибка1\" не было")
	}
}
