package main

import (
	"fmt"
	"reflect"
)

type MyType struct {
	Name string `json:"name"`
}

func main() {
	mt := MyType{Name: "set"}
	mtType := reflect.TypeOf(mt) //打印的是变量类型
	fmt.Println(mtType)
	name := mtType.Field(0) //name这个结构体
	fmt.Println(name)
	tag := name.Tag.Get("json") //结构体json tag
	fmt.Println(tag)
	value := name.Name //定义结构体的名称
	fmt.Println(value)

}
