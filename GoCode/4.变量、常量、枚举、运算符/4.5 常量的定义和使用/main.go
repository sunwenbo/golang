package main

import "fmt"

func main() {
	//常量的定义和使用
	//常量的位置是在数据区
	//变量的位置在栈区，系统为每一个应用程序分配1M空间来存储变量 在程序运行结束后会自动释放
	var s1 int = 10
	var s2 int = 20
	//常量一般用大写字母用于表示
	const A int = 10
	//a = 20  //常量的值不允许修改
	//常量的内存区不可以通过& 取地址来访问
	fmt.Println(A)
	fmt.Println(&s1,&s2)
	//字面常量
	fmt.Println("hello word")
	//硬常量 32
	s3 := A + 32
	fmt.Print(s3)
}