package main

import (
	"encoding/json"
	"fmt"
)

type human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func parseJsonObject() {
	var obj interface{}
	var h = human{
		Name: "name",
		Age:  1,
	}
	humanStr, err := json.Marshal(h)
	if err != nil {
		fmt.Println("错误")
	}
	err = json.Unmarshal(humanStr, &obj)
	objMap, _ := obj.(map[string]interface{})
	for k, v := range objMap {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string,value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{},value is %v\n", k, value)
		default:
			fmt.Printf("type of %s is unknow,value is %v\n", k, value)
		}
	}
}
func main() {
	parseJsonObject()
}
