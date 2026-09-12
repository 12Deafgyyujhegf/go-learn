package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("sb")
	case age >= 18:
		fmt.Println("bigsb")
	default:
		fmt.Println("no")
	}

	var week int
	fmt.Println("星期")
	fmt.Scan(&week)

	switch week {
	case 1:
		fmt.Println("one")
		fallthrough //顺次往下
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("jdsk")
	}
}
