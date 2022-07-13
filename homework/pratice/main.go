package main

import "fmt"

/**
 * Author: Leo Brown
 * Description: 将修改字符串
 */
func main() {
	strArray1 := []string{"I","am","stupid","and","weak"}
	for index,value := range strArray1{
		if value == "stupid"{
			strArray1[index] = "smart"
		}
		if value == "weak"{
			strArray1[index] = "strong"
		}
	}
	fmt.Printf("输出的内容:%+v",strArray1)
}
