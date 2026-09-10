package main

import "fmt"

func main() {

	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)
	if age <= 18 {
		if age <= 0 {
			fmt.Println("未出生")
		} else {
			fmt.Println("少年")
		}
	} else {
		if age >= 50 {
			fmt.Println("老年")
		} else {
			fmt.Println("成年")
		}
	} //else必须和“}”同一行
}
