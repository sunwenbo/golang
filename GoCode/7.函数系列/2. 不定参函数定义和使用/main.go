package main

import "fmt"

//不定参函数 函数参数类型为 ...数据类型
func test(args ...int) {
	//fmt.Println(len(args))
	//len() 计算字符串有效长度 计算不定参函数的长度
	//for i := 0; i < len(args);i++ {
	//	fmt.Println("下标",i,"值",args[i])
	//}
	sum := 0
	// for 和range 可以遍历 集合中的数据信息，数组切片 结构体数组 map
	//第一个参数为下标，第二个为值
	for _,data := range args{  //匿名参数
		sum += data
	}
	fmt.Println(sum)
}

func main()  {
	test(1,2,3)
	test(1,2,3,4,5)
}