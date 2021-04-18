package main

import "fmt"

//复合类型主要包括了数组、指针、切片、结构体等

func main01() {
	//数组的定义和使用  数组是在内存中连续存储的多个相关类型的数据集合
	//var 数组名 [元素个数]数据类型
	//数组下标是从0开始的，到数组下标最大值为数组元素个数-1
	var arr[10] int = [10]int{1,2,3,4,5,6,7,8,9,10}
	arr[0] = 123
	arr[1] = 234
	arr[2] = 345
	//数组名[下标]=值
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
	fmt.Println(arr)

	for i := 0; i < 10; i++ {
		fmt.Println(arr[i])
	}

	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//range 遍历一个集合的数据
	for i,v := range arr {
		fmt.Println(i,v)
	}
}

func main02() {
	//数据在定义时，可以初始化部分元素的值
	//var arr[10] int = [10] int {1,2,3,4,5}
	//使用自动类型推导创建数组
	//arr := [10]int{1,2,3,4,5}
	arr := [...]int{1,2,3}
	for i,v := range arr {
		fmt.Println(i,v)
	}
	fmt.Printf("%T",arr)
}

func main() {
	//在定义时 叫元素个数，
	arr := [5]int{1,2,3,4,5}
	//在数组使用时叫下标
	//arr[5] = 100  数组下标越界
	//数组在定义后，元素个数不可改变  数组是一个常量， 不允许赋值 数组名代表整个数组
	fmt.Println(arr)
	//数组名也可以表示数组在内存中的首地址
	fmt.Printf("%p\n",&arr)
	fmt.Printf("%p\n",&arr[0])
}
