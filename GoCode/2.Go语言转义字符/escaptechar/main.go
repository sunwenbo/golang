package main

import "fmt" //fmt包中提供格式化，输出，输入的函数

func main() {
	//演示转义字符的使用
	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("/Users/sunwenbo/go/GoCode/2.Go语言转义字符/escaptechar")
	fmt.Println("tom 说\"i love you\"")
	// \r 回车，从当前行的最前面开始输出，覆盖掉以前的内容，\r后面有几个字符就会覆盖掉几个字符
	fmt.Println("天龙八部雪山飞狐\r张飞")
}
