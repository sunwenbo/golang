package main

import "fmt"

func test(arr [10]int) {
	for i := 0; i < len(arr);i++ {
		fmt.Println(i)
	}
	arr[2] = 123
	fmt.Printf("%p\n",&arr)
}

func main01() {
	arr := [...]int{1,4,3,6,8,2,9,0,5,7}
	//fmt.Println(arr)
	//数组作为函数参数传递数组名
	//数组作为函数参数是作为值传递的 形参不会影响实参
	test(arr)
	fmt.Println(arr)
	fmt.Printf("%p",&arr)
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

func main() {
	arr := [10]int{1,4,3,6,8,2,9,0,5,7}
	//只读变量
	arr = BubbleSort(arr)
	fmt.Println(arr)
}