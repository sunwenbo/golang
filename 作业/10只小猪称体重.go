package main

import "fmt"

func main()  {
	arr := [...]int{1,3,4,6,2,7,8,10,9,5,11}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println("最重的小猪体重是:",max )
}
