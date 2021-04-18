package main

import (
	"fmt"
	"unsafe"
)

//演示golang中整数类型使用
func main(){
	var i int = 1
	fmt.Println("i=",i)

	//测试一下int8的范围 -128～127
	//其它的 int16,int32,int64，类推....
	var j int = -129
	fmt.Println("j=",j)

	//测试一下uint8的范围 (0~255) 其它的 uint16,uint32,uint64，类推....
	var k uint8 = 255 //不能超过-1  和256
	fmt.Println("k=",k)

	//int、uint、rune、byte的使用
	var a int = 8900
	fmt.Println("a=",a)
	var b int = 1
	fmt.Println("b=",b)
	//byte 等价于 uint8
	var c byte = 255
	var d byte = 1
	fmt.Println("c=",c,"d=",d)

	//整型的使用细节
	var n1 = 1000 // n1 是什么类型？
	//这里我们给介绍一下如何查看某个变量的数据类型
	//fmt.Printf() 可以用于做格式化输出
	fmt.Printf("n1的类型为%T",n1)

	//如何在程序查看某个变量的字节大小和数据类型(使用较多)
	var n2 int64 = 10
	//unsafe.Sizeof(n2) 是unsafe包的一个函数，可以返回n2变量占用的字节数
	fmt.Printf("n2的类型为%T, n2占用的字节数是 %d",n2, unsafe.Sizeof(n2))

}