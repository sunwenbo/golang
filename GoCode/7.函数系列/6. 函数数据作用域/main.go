package main

import "fmt"

//func demo1() {
//	fmt.Println(a)
//}

func main01() {
	a := 10
	//demo1()
	//局部变量 在函数内部定义的变量 作用域限定于本函数内部 从变量定义到本函数结束有效
	//var a 变量名 数据变量 先定义后使用
	//在同一作用域范围内 变量名是唯一的

	//匿名内部函数
	{
		var a int = 20   //定义一个内部变量a值为20
		//a = 20   //给a重复赋值
		fmt.Println(a)
	}
	fmt.Println(a)

	//程序中如果出现了相同的变量名，如果本函数有自己的变量就使用自己的，
	//如果没有，就上外层寻找变量  如果名字相同会采用就近原则
	i := 0
	for i := 0 ; i < 10; i++ {

	}
	fmt.Println(i)
}

//全局变量 在函数外部定义的变量称为全局变量
//全局变量作用域是项目中所有文件
//全局变量在内存中数据区村粗 和const定义的常量存储位置都是数据区

var a int = 10  //如果变量没有定义初始值则为0 int类型
const b int = 10  //全局常量
func main() {
	demo2()
	//在局部变量作用范围内，全局变量不起作用，采用就近原则
	//a = 234
	fmt.Println(a)
}

func demo2()  {
	//a = 123
	fmt.Println(&a)
}