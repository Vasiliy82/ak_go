package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	// добавил эту строку
	ch = nil

	time.Sleep(50 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}

/*

Задача №6 «Нострадамус». В задаче «Болтушка» запускались
горутины…В рамках задачи необходимо посмотреть внимательно на
звёзды и сказать: сколько горутин запускалось, а также рассказать о
дальнейшей их судьбе.

Запускалась главная рутина + 2шт горутины. Первая содержит в себе антипаттерн "холостой цикл" (из-за наличия default
внутри select и отсутствия задержки цикл будет повторяться бесконечное количество раз)

*/
