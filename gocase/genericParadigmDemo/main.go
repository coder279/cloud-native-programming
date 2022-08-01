package main

import "fmt"

type Order interface{
~int8 | ~int64
}

func smallest[T Order](s T) T {
	fmt.Println(s)
	return s
}

func main() {
	var a int64 = 32
	smallest(a)
}
