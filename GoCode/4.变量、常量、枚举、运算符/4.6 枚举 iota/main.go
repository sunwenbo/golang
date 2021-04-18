package main

import (
	"fmt"
)

func main() {
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)
	fmt.Println(a,b,c)
	//定规变量为状态
	status := a
	println(status)
	status=b
	println(status)
	status=c
	println(status)

	const (
		a1 = iota //自动增长
		b1
		c1
		d1
	)
	fmt.Println(a1,b1,c1,d1)

	const (
		//如果定义枚举时常量写在同一行，值相同。换一行值加1
		//在定义枚举时可以为期赋初始值，但是换行后不会自增长
		a2 = iota
		b2,c2 = iota,iota
		d2,e2
	)
	fmt.Println(a2,b2,c2,d2,e2)
}