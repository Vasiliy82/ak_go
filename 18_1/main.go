/*Задача №1. Необходимо код примера 1 изменить так, чтобы tcp-сервер обрабатывал подключения параллельно.*/
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	// корректное завершение работы по ctrl+c
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		<-sigs
		fmt.Println("Ctrl+C was pressed, shutdown...")
		// чтобы прервать блокирующий вызов в основном цикле
		listener.Close()
		cancel()
	}()

OUT:
	for {
		select {
		case <-ctx.Done():
			break OUT
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleConn(ctx, conn)
		}
	}
	log.Println("завершение работы")
}
func handleConn(ctx context.Context, c net.Conn) {
	defer c.Close()
	var i int
	for {
		select {
		case <-ctx.Done():
			io.WriteString(c, "Server go down, bye")
			return
		default:
			_, err := io.WriteString(c, fmt.Sprintf("%d\n", i))
			if err != nil {
				log.Println(err)
				return
			}
		}
		i++
		time.Sleep(time.Second)
	}
}
