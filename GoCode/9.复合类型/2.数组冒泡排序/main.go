package main

import "fmt"

func main() {
	arr := [...]int{1,4,3,6,8,2,9,0,5,7}
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//比较两个相邻的元素 满足条件交换数据
			//升序使用大于号， 降序使用小于号
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}