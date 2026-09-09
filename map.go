package main

import "fmt"

func main() {

	var name = map[int]string{
		1: "我是",
		2: "你爹",
		3: "",
	}
	fmt.Println(name)
	fmt.Printf("%#v\n", name[3])
	fmt.Println(name[4])
	jia := name[4]
	fmt.Println("单值取 name[4]:", jia)

	jia, ok := name[4]
	fmt.Println("双值取 name[4]:", jia, ok)

	jia2, ok2 := name[3]
	fmt.Println("双值取 name[3]:", jia2, ok2)

	name[1] = "我就是"
	fmt.Println(name)
	delete(name, 3)
	fmt.Println(name)
}
