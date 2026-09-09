package main

import (
	"fmt"
	"sort"
)

func main() {
	var clist []string
	clist = append(clist, "你好")
	clist = append(clist, "我是")
	clist = append(clist, "嫩爹")
	fmt.Println(clist)
	fmt.Println(clist == nil)

	var alist []string
	fmt.Println(alist == nil)
	blist := make([]int, 3) //既可以使blist不为空数组，又能给这个数组全部赋值为零
	fmt.Println(blist == nil)
	fmt.Println(blist)

	var dlist = []int{34, 12, 54, 15, 423}
	sort.Ints(dlist) //升序
	fmt.Println(dlist)
	sort.Sort(sort.Reverse(sort.IntSlice(dlist))) //降序
	fmt.Println(dlist)
}
