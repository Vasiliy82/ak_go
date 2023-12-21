package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Привет из дочерней горутины")

	}()

	time.Sleep(time.Second)

	ch1 := make(chan string)
	go func() {
		ch1 <- "Привет, строковой канал"

	}()
	fmt.Println(<-ch1)

	time.Sleep(time.Second)

	ch2 := make(chan string, 4)
	ch2 <- "Привет!"
	ch2 <- "Буферизованный канал"
	for len(ch2) > 0 {
		fmt.Println(<-ch2)
	}

}
