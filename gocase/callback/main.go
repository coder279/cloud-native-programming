package main

import "fmt"

func add() (func(x, y int), error) {
	return func(x, y int) {
		fmt.Println(x + y)
	}, nil
}

func increase(x, y int) {
	fmt.Println(x + y)
}
func main() {
	f,err := add()
	if err != nil{
		fmt.Println("错误")
	}
	f(1,2)
}
