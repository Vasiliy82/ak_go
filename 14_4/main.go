package main

/*
Задача №5. «Айболит: Гнев Милосердия»
В рамках задачи будем работать с картотекой известного врача.
Нужно будет написать модуль с несколькими версиями: v1.0.0, v1.1.0,
v2.0.0, v2.1.0. Модуль должен прочитать файл со следующим
содержимым:
{"name":"Ёжик","age":10,"email":"ezh@mail.ru"}
{"name":"Зайчик","age":2,"email":"zayac@mail.ru"}
{"name":"Лисичка","age":3,"email":"alice@mail.ru"}
*/
import (
	format "github.com/Vasiliy82/ak_go/format/v2"
)

func main() {

	format.Do("input", "output")

}
