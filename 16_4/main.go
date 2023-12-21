/*
Задача №4.Необходимо создать функцию start, которая в консоль
выводит некоторое сообщение. Необходимо запустить 10 горутин,
которые будут запускать функцию start и выводить в консоль факт
своего запуска, причём необходимо обеспечить однократный запуск
функции start.
*/
package main

import (
	"fmt"
	"sync"
)

var runOnce sync.Once

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("routine ", i, "started")
			runOnce.Do(start)
			fmt.Println("routine ", i, "done")
		}()
	}

	wg.Wait()

}

func start() {
	fmt.Println("Start()")
}
