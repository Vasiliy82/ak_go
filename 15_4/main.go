/*
Задача №4 «Цаца». Если запустить следующий код программы:
возникнет ошибка блокировки:
fatal error: all goroutines are asleep - deadlock!

Нужно дописать6 логику так, чтобы ошибки блокировки не возникало.
Вместо ошибки в консоль должна выводится фраза «happy end». Так
как канал ch является «особым»7, то в него ЗАПРЕЩАЕТСЯ писать
значения.

*/

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{})

	go func() {
		<-ch
		stop <- struct{}{}
	}()
	// добавил эту строку
	close(ch)
	<-stop
	fmt.Println("happy end")
}
