package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := func3(); err != nil {
		if innerErr := errors.Unwrap(err); innerErr != nil {
			// в задаче просят вывести эту ошибку
			fmt.Println(innerErr)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Всё хорошо")
	}

}

func func1() error {
	return errors.New("ошибка1")
}

func func2() error {
	if err := func1(); err != nil {
		return fmt.Errorf("ошибка2:%w", err)
	}
	return nil
}

func func3() error {
	if err := func2(); err != nil {
		return fmt.Errorf("ошибка3:%w", err)
	}
	return nil
}
