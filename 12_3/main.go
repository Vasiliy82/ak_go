package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("ошибка1")

	// кейс, если ошибка была
	err2 := fmt.Errorf("ошибка2:%w", err1)
	// кейс, если не было
	// err2 := fmt.Errorf("ошибка2:%w", errors.New("ошибка1"))
	err3 := fmt.Errorf("ошибка3:%w", err2)

	fmt.Println(err3)

	// Не используя unwrap, нужно определить была ли ошибка «ошибка1»
	if found := errors.Is(err3, err1); found == true {
		fmt.Println("Вложенная ошибка \"ошибка1\" была")
	} else {
		fmt.Println("Вложенной ошибки \"ошибка1\" не было")
	}

}
