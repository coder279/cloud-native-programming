package main

import (
	_ "cloud-native-programming/gocase/init/a"
	_ "cloud-native-programming/gocase/init/b"
	"fmt"
)

func init() {
	fmt.Println("init main")
}

func main() {

}
