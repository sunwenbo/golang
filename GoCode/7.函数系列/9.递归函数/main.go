package main

import "fmt"

//在函数定义时调用函数本身叫做递归函数
//死递归
func test(a int) {
	//在递归函数中一定要有出口  return
	if a == 0 {
		return
	}
	a--
	test(a)
	fmt.Println(a)
}

func main01() {
	test(10)
}


//计算N的阶乘
var sum int = 1

func test1(n int) {
	if n == 1 {
		return
	}
	test1(n-1)
	sum *= n
}

func main() {
	test1(3)
	fmt.Println(sum)
}