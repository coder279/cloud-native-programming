package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	done := make(chan bool)
	defer close(message)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process done")
				return
			default:
				fmt.Printf("send message %d\n", <-message)
			}
		}
	}()
	for i := 0; i < 10; i++ {
		message <- i
	}
	time.Sleep(5 * time.Second)
	close(done) //close操作相当于向done赋值
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit")
}
