package main

import "fmt"

func main() {
	a := 10
	b := 20
	//匿名内部函数，在函数中间再定义一个函数
	//f 是函数类型对应的变量
	f := func (a int, b int) {
		fmt.Println(a+b)
	}
	f(a,b)
	f(20,30)

	func (){
		fmt.Println("hello word")
	}()

	a1 := 10
	b1 := 20
	f1 := func (a int,b int) int{
		return a+b
	}
	//函数调用 f 函数类型 v 接收返回值 类型为int
	fmt.Printf("%T\n",f1)
	v := f1(a1,b1)
	fmt.Printf("%T\n",v)
	fmt.Println(f1(a1,b1))
}
