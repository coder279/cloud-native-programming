package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int, 10)
	ticker := time.NewTicker(1 * time.Second)
	for{
		go func() {
			for _ = range ticker.C {
				queue <- 1
				fmt.Println("生产了一条数据")
			}
		}()
		go func() {
			for _ = range ticker.C {
				fmt.Printf("消费了一条%d\n", <-queue)
			}
		}()
	}

}
