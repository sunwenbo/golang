package main

import "fmt"

// func 函数名 （形参列表）返回值类型列表{代码体}

//第一种
func sub(a int ,b int) int{
	sum := a - b
	//表示函数的结束 return 后面的代码不会执行 同时return也会将函数的返回值传递给主调函数
	return sum
}

//第二种
func sub1(a int ,b int) int{
	//return 可以写一个表达式，将表达式的结果作为返回值
	return a - b
}
//第三种
func sub2(a int ,b int) (sum2 int) {
	//return 可以写一个表达式，将表达式的结果作为返回值
	sum2 = a -b
	return
}

//函数多个返回值
//如果函数没有参数在函数调用时()必须写
func test() (a int ,b int, c int) {
	//多重赋值
	a,b,c = 1,2,3
	return
}


func main()  {
	//value 接收返回值
	value1 := sub(20,10)
	fmt.Println(value1)
	value2 := sub1(30,10)
	fmt.Println(value2)
	value3 := sub2(40,10)
	fmt.Println(value3)

	//函数的返回值可以为主调函数进行赋值 可以通过返回值改变实参数据
	a,b,c := test()
	fmt.Println(a,b,c)
	//如果函数有多个返回值，但是不使用可以用匿名变量'_'来实现不接收
	c,_,d := test()
	fmt.Println(c,d)
}


