package main

import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var name3 = "jack"

//上面的声明方式，也可以改成一次性声明
var (
	n3 = 300
	n4 = 900
	name4 = "mary"
)

func main(){
	//该案例演示golang如何一次性声明多个变量
	//var n1,n2,n3 int
	//fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
	////一次性声明多个变量的方式2
	//var n4, name, n5 = 100,"tom",888
	//fmt.Println("n4=",n4,"name=",name,"n5=",n5)
	////一次性声明多个变量的方式3,同样可以使用类型推导
	//n6,name1,n7 := 100,"tom2",888
	//fmt.Println("n6=",n6,"name1=",name1,"n7=",n7)
	//输出全局变量
	fmt.Println("n1=",n1,"n2=",n2,"name=",name3)
	fmt.Println("n3=",n3,"n4=",n4,"name4=",name4)
	//交换变量的值
	a := 10
	b := 20
	a,b = b,a
	fmt.Println(a,b)
	//匿名变量
	_,a1,_,b1 := 1,2,3,4
	fmt.Println(a1,b1)
}
