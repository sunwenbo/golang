package main

import (
	"fmt"
)

//不定参函数 函数参数类型为 ...数据类型
func test2(a int , b int ) {
	sum := a + b
	fmt.Println(sum)
}
func test1(a int , b int ) {
	test2(a,b)
}

func test3(args ...int)  {
	// 如果函数的参数为不定参  传递方式为args[0:]...  0为下标 :后面不写代表最大下标
	test4(args[0:]...)
}

func test4(args ...int)  {
	for i :=0; i < len(args);i++ {
		fmt.Println(i,args[i])
	}
}

//所有的函数都是全局函数，可以被项目中所有的文件使用 在项目中函数名是唯一的
func main()  {
	test1(10,20)
	test2(20,20)
	test3(1,2,3,4)
}
