package main

import "fmt"

func test() {
	fmt.Println("瓜娃子")
}

func test2() {
	fmt.Println("细伢子")
}

func test1(a int,b int){
	fmt.Println(a+b)
}

func test3(a int,b int) int {
	fmt.Println(a+b)
	return a+b
}

func test4(a int,b int)  {
	fmt.Println(a+b)
}

func test5() {
	test4(10,20)
}


//type 可以定义函数类型
//type 可以为已存在类型起别名
type FUNCTYPE func()
type FUNCTEST func(int,int)
type TypeSum func(int,int) int

func main() {
	//定义函数类型变量
	//函数类型其实就是一个指针
	var f FUNCTYPE
	f = test
	//通过函数变量调用函数
	f()

	var f1 FUNCTYPE
	f1 = test2
	f1()

	var f2 FUNCTEST
	f2 = test1
	f2(10,20)

	var f3 TypeSum
	f3 = test3
	f3(30,10)

	//如果使用print打印函数名是一个地址
	//函数名本身就是一个指针类型数据  在内存中的代码区存储
	fmt.Println(test4)

	//自动类型推导创建函数类型
	f4 := test4
	f4(30,40)
	//打印f4类型
	fmt.Printf("%T\n",f4)

	var f5 func(int,int)
	f5 = test4
	f5(10,20)
}