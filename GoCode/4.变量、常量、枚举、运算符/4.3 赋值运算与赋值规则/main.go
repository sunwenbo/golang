package main
import "fmt"

//变量使用的注意事项
func main(){
    //该区域的数据值可以在同一个类型范围内不断比变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=",i)
	//i = 1.2   不能改变值的数据类型
	//fmt.Println("i=",i)

	//var i int = 50  //在同一个函数中不允许重复定义
	var a = 1
	var b = 2
	var c = a + b
	fmt.Println("c=",c)
	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2
	fmt.Println("res=",res)
}