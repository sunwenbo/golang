package main

import (
	"fmt"
	"math/rand"
	"time"
)
func BubbleSort(arr [6]int)  [6]int{
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
	var red [6]int
	var blue [1]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(red); i++{
		//遍历之前存在的值和新随机数是否有重复
		temp := rand.Intn(33)+1
		for j := 0; j < i; j ++{
			if temp == red[j] {
				temp = rand.Intn(33)+1
				j = -1
				continue
			}
		}
		red[i] = temp
	}

	for i := 0; i < len(blue); i++{
		blue[i] = rand.Intn(16)+1
	}

	//red = BubbleSort(red)
	fmt.Printf("红球为%d,篮球为%d",red,blue)
}