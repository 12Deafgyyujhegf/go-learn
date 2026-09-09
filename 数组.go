package main

import "fmt"

func main() {

	var list [3]string = [3]string{"你好", "我只", "帅"}
	fmt.Println(list)
	list[0] = "这些话"
	fmt.Println(list)
	fmt.Println(len(list))
}
