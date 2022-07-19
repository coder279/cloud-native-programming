package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b") //通过context进行值传递
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second) //context的timeout
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt")
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit")
	}
}
