package main

import "fmt"

func main() {

	name := "胡"
	fmt.Println(name)
	var aa byte = 'a'
	fmt.Printf("%d %c\n", aa, aa)
	var bb rune = '中'
	fmt.Printf("%d %c\n", bb, bb)
	fmt.Printf("%T %T\n", "你好", 4)
	fmt.Printf("%v\n", "工作")
	fmt.Printf("%v\n", 909)
	fmt.Printf("%#v", "")

}
