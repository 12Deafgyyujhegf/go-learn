package main

import "fmt"

func he(numberlist ...int) { //传不定参数
	var sum int
	for _, i1 := range numberlist { //range是一个便利函数
		sum += i1
	}
	fmt.Println(sum)
}

func main() {
	// he(1, 2, 3)
	// he(3, 4, 5)
	// a, b := r3() //承接返回值
	// fmt.Println(a, b)

	// r5()
	var index int
	fmt.Scan(&index)
	switch index {
	case 1:
		login()
	case 2:
		register()
	case 3:
		usercenter()
	}

}

func r1() bool {
	return true
}

func r2() (a string, b bool) {
	return a, b
}
func r3() (a string, b bool) {
	a = "123"
	b = true
	return a, b
}

//匿名函数，go语言里函数内不能再创建函数，可以使用匿名函数

func r5() {
	var getname = func() string {
		return "我你爹"
	}
	var setname = func(name string) { //是等于号呦
		fmt.Println(name)
		return
	}
	setname("张三")
	fmt.Println(getname())
}

func login() {
	fmt.Println("登录")
}

func register() {
	fmt.Println("注册")
}

func usercenter() {
	fmt.Println("用户中心")
}
