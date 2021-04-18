package main

import "fmt"

func main1() {
	//声明变量
	//var a int
	//&c 运算符 取地址运算符
	//fmt.Scan(&a)
	//%p 占位符 表述输出一个数据对应的内存地址 ¥&a
	//0x表示十六进制数据
	//age :=20
	//fmt.Printf("age:%d\n", age)
	//fmt.Printf("a=%d\n",a)
	//fmt.Printf("%p",&a)
	//空格或候车作为接受结束
	//var str string
	//fmt.Scan(&str)
	//fmt.Println(str)
	//fmt.Printf("%p",&str)
	//接收两个数据
	var s1,s2 string
	fmt.Scan(&s1,&s2)
	//fmt.Scanf
	fmt.Println(s1,s2)
}
func main() {
	var user string
	var passwd string
	fmt.Println("请输入用户名：")
	fmt.Scanf("%s",&user)
	fmt.Println("请输入密码：")
	fmt.Scanf("%s",&passwd)
	fmt.Println(user,passwd)

}
