/*Задача №3. Нужно запустить 5 горутин и остановить в некоторое
время, которое рассчитывается по формуле: текущий момент + 2
секунды.*/

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
	deadline := time.Now().Add(2 * time.Second)
	myCtx, cancel := context.WithDeadline(ctx, deadline)
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
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время выполнения:", duration)
}
