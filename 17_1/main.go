/*Задача №1. Нужно запустить 5 горутин и остановить, используя
контекст с отменой.*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	myCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		i := i
		wg.Add(1)
		go func(innerCtx context.Context) {
			defer wg.Done()

			for {
				select {
				case <-innerCtx.Done():
					fmt.Println("Routine ", i, " done, reason: ", innerCtx.Err())
					return
				default:
					fmt.Println("Routine ", i, " impact")
					time.Sleep(time.Second)
				}
			}

		}(myCtx)
	}
	cancel()
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время выполнения:", duration)
}
