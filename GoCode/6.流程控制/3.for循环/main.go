package main

import "fmt"

/*
语法结构如下：
for 表达式1;表达式2;表达式3{
	循环体
}
表达式1：定义一个循环的变量，记录循环的次数
表达式2：一般为循环条件，循环多少次
表达式3：一般为改变循环条件的代码，使循环终有一天不再成立
循环体：重复要做的事情
*/

func main01() {
	for a := 1; a <= 10; a++ {
		fmt.Println(a)
	}
}

////1-100的和
//func main02() {
//	sum := 0
//	for i := 1; i <= 100; i++ {
//		sum += i
//	}
//	fmt.println(sum)
//}

//计算100以内偶数的和
//func main() {
//	sum := 0
//	for i := 1; i <= 100; i++{
//		if i % 2 == 0 {
//			sum += i
//		}
//	}
//	fmt.println(sum)
//}
//func main() {
//	sum := 0
//	for i := 0; i <= 100; i+=2{
//		sum += i
//	}
//	fmt.println(sum)
//}
func main02() {
	//for函数格式，i是在函数内部定义的不能在外部使用
	//局部变量
	//for a := 1; a <= 10; a++ {
	//	fmt.Println(a)
	//}
	//fmt.Println(i)
	//i := 1
	//for ;i <= 10; i++{
	//	fmt.Println(i)
	//}
	//fmt.Println(i)
	i := 1
	//表达式可以写在下面
	for ;; {
		if (i>=10){
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println(i)
	//死循环 条件永远为true 在程序中要避免死循环的出现
	//for {
	//	fmt.Println("hello word")
	//  break 使用break跳出循环
	//}
}

//水仙花 一个三位数100-999 各个位数的立方的和等于这个数本身 就是一个水仙花数

func main03() {
	i := 1
	for ; i < 1000;i++ {
		a := i / 100 //百位
		b := i / 10 % 10 //十位
		c := i % 10 //个位
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}


// 1- 100 逢7 含有7的 7的倍数敲桌子
func main() {
	for i := 1; i<= 100; i++ {
		if i % 7 == 0 || i % 10 == 7 || i / 10 ==7 {
			fmt.Println("敲桌子")
		} else {
			fmt.Println(i)
		}
	}
}






