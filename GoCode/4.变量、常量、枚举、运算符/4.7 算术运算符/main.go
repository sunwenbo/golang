package main

import "fmt"

// 运算运算符 + - * / % ++ -- += -= *= /=
// 关系运算符 == > < >= <= !=  用于程序中的逻辑判断，返回值为布尔值
// 逻辑运算符 取反!  &&逻辑与  ||逻辑或（逻辑与高于逻辑或）  返回值为布尔值
// 取值运算符 *取指针变量 &取回内存地址
//除法运算符
func main01() {
	a := 10
	b := 20
	// 整型数据相除结果为整型 保留整数部分为0 被除数不能为0
	c := a / b
	fmt.Println(c)
}
//取余运算符

func main02() {
	a := 10
	b := 3
	// 除数不能为0
	// 取余运算只能用于整数数据
	c := a % b
	fmt.Println(c)
}
//自增自减运算符
func main03() {
	a := 10
	// 不能将自增自减运行在表达式中  ++ -- 只能写在变量的后面，叫做后自增，后自减
	a++  //自增运算符
	a--  //自减运算符
	fmt.Println(a)
}

func main04()  {
	a := 46
	b := 7
	c := a / b
	d := a % b
	fmt.Printf("46天包含%d周和%d天",c,d)

}

func main05()  {
	//编程实现107653秒是几天几小时几分钟几秒？
	s := 107653
	fmt.Println("天",s/60/60/24%365)
	fmt.Println("时",s/60/60%24)
	fmt.Println("分",s/60%60)
	fmt.Println("秒",s%60)
}

func main() {
	var year int
	fmt.Println("请输入一个年份:")
	fmt.Scan(&year)
	//能够被400整除，2. 能够被4整除 不能被100整除

	b := (year % 400 == 0) || (year % 4 == 0 && year % 100 != 0)
	fmt.Println(b)
}