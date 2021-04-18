package main

import (
	"fmt"
)

func main01() {
	//数组定义 var 数组名 [元素个数]  数据类型
	//切片定义 var 切片名 []  数据类型
	//var s []int
	//fmt.Println(s)
	//自动推导类型创建切片 make([]数据类型,5)
	s := make([]int,5)
	s[0] = 123
	s[1] = 234
	s[2] = 345
	s[3] = 456
	s[4] = 567
	//s[5] = 678    //err
	//通过append  添加切片元素
	s = append(s,678,789,8910)

	fmt.Println(s)

	//通过len查看切片长度
	fmt.Println(len(s))
}

func main02() {
	s := make([]int,5)
	s[0] = 123
	s[1] = 234
	s[2] = 345
	s[3] = 456
	s[4] = 567

	//遍历
	for i := 0 ; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for i , v := range  s{
		fmt.Println(i,v)
	}
}

func main03() {
	//不写元素个数叫切片， 数组必须写元素个数。 切片可以扩展
	//切片定义
	//var s  []int = []int{1,2,3,4,5}
	s := []int{1,2,3,4,5,}
	fmt.Println(s)
	fmt.Printf("s的类型为%T\n",s)
	s = append(s,7,8,9)
	fmt.Println(s)


	//数组定义
	arr := [...]int{2,3,1,4,5,9,11,7,6,8}
	fmt.Println(arr)
	fmt.Printf("arr的类型为%T\n",arr)

	//容量要大于等于长度
	var s1 []int = []int{1,2,3,4,5}
	s1 = append(s1,6,7,8,9)
	fmt.Println("长度",len(s1))
	fmt.Println("容量",cap(s1))
	//容量每次扩展为上次的倍数
	s1 = append(s1,6,7,8,9)
	fmt.Println("长度",len(s1))
	fmt.Println("容量",cap(s1))
	//如果整体数据没有超过1024字节，每次扩展为上一次的倍数，超过1024 每次扩展上次的1/4
	s1 = append(s1,6,7,8,9)
	fmt.Println("长度",len(s1))
	fmt.Println("容量",cap(s1))

}

func main()  {
	main03()
	s := make([]int,5)
	s[0] = 1
	//使用append在长度后添加数据，同时增加容量
	s = append(s,1,2,3)
	fmt.Println(s)

	//在切片打印时，只能打印有效长度中的数据  cap不能作为数据长度打印的条件
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

