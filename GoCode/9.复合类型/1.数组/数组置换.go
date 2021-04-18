package main

import "fmt"

func main() {
	arr := [...]int{2,3,1,4,5,9,11,7,6,8}
	start := 0
	end := len(arr) - 1

	for {
		if (start > end) {
			break
		}
		arr[start],arr[end]=arr[end],arr[start]
		start ++
		end --
	}
	fmt.Println(arr)
}
