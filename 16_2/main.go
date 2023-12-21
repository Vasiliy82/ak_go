/*Задача №2. Используя пакет atomic, необходимо реализовать счётчик,
с которым параллельно могут работать несколько горутин.*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ctrAsyncInt32 struct {
	i int32
}

func (c *ctrAsyncInt32) IncrementAsync() {
	atomic.AddInt32(&c.i, 1)
}
func (c *ctrAsyncInt32) GetAsync() int32 {
	return atomic.LoadInt32(&c.i)
}

func main() {
	c := ctrAsyncInt32{i: 0}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			// множество горутин, каждая из которых максимально быстро будет инкрементировать значение счетчика
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.IncrementAsync()
			}
			fmt.Println("routine incrementer ", i, "finished")
		}()
	}

	wg.Wait()

	fmt.Println("Final counter value is: ", c.GetAsync(), " (must be 100 * 1000 = 100.000)")

}
