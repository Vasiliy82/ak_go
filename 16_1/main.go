/*
Задача №1. Необходимо запустить 5 горутин. Не используя time.Sleep
нужно обеспечить вывод в консоль каждой горутиной своего
уникального сообщения. Например:
горутина: 1
горутина: 2
…
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Горутина ", i)
		}()
	}
	wg.Wait()
}
