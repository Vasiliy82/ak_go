package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d Duck) Sing() string {
	return d.voice
}
func main() {
	var d *Duck
	// проблематика: в том, что интерфейс - это не совсем указатель: внутри содержится тип и указатель
	// как решить:
	// Вариант №1 - вынести проверку d на nil в main
	// Вариант №2 - в функции func (d Duck) Sing() обращением к полю (return d.voice) проверить его наличие
	// Вариант №3 - избегать вообще избегать использования nil совместно с интерфейсами
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}
func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
