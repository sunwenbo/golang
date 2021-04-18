package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main01() {
	//1、导入头文件 math/rand   time
	//2、添加随机数种子
	//3、使用随机数
	start := time.Now().Nanosecond()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 10; i++ {
		//根据时间 1970.1.1.0.0.0 秒
		fmt.Println(rand.Intn(123))
	}
	fmt.Println(time.Now())
	fmt.Println(time.Now().Second())
	end := time.Now().Nanosecond()
	fmt.Println(end - start)
}

func BubbleSort(arr [10]int)  [10]int{
	for i :=0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if (arr[j] > arr[j+1]) {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}

func Sort(arr [10]int )  [10]int{
	for i := 0 ;i < len(arr)-1; i++{
		for j := 0; j < len(arr)-1-i; i++{
			if (arr[j] > arr[j+1]) {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return  arr
}

func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i :=0; i< len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	arr  = Sort(arr)
	fmt.Println(arr)
}









