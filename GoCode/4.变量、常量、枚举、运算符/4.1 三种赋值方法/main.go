package main

import (
	"fmt"
)

func main1(){
	//golang的变量使用方式1
	//第一种：指定变量类型，声明后若不赋值，使用默认值
	// int 的默认值就是0，其他数据类型的默认值后面介绍
	var i int
	fmt.Println("i=",i)
	//第二种：根据值自行判定变量类型（类型推导）
	var num = 10.11
	fmt.Println("num=",num)
	//第三种：省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	//自动推导类型 变量名 := 值 会根据变量选择类型
	//下面的方式等价于 var name string name = "tom"
	//变量的类型不通不能进行计算，需要转换为相同类型
	name := "tom"
	fmt.Println("name=",name)
}
func main(){
	//计算园的面积和周长
	//面积 PI * r * r
	//周长 2 * PI * r
	//var PI float64=3.14159
	PI := 3.14159
	r := 2.5
	//var s float64
	//var l float64
	//计算面积
	s := PI * r * r
	//计算周长
	l := 2 * PI * r
	println("面积:" ,s)
	println("周长:", l)
}
