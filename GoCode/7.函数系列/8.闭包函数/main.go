package main

import "fmt"

//函数在调用结束会从内存销毁
func test(a int) {
	a++
	fmt.Println(a)
}
func main01() {
	a := 0
	for i := 0; i <=10;i++ {
		test(a)
	}
}

//可以通过匿名函数和闭包 实现函数和栈区的本地化操作
func test1() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}


func main() {
	//不加括号是将test1函数类型赋值给f
	//加了括号将test1的返回值给f
	f := test1()  //返回函数类型
	//f := test   //返回匿名函数类型
	fmt.Printf("%T",f)
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}


}