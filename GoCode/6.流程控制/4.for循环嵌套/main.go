package main

import "fmt"

//func main() {
//	//嵌套循环中，执行次数位外层 * 内层
//	 for i := 0; i <5; i++ {
//		for j :=0; j <5; j++ {
//			fmt.Println(i,j)
//		}
//	 }
//}
func main() {
	//外层控制行，内层控制列
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d= %d\t",j ,i,(i * j))
		}
		fmt.Println("")
	}
	for a := 1; a <= 5; a++ {
		if a == 4 {
			continue
		}
		fmt.Println(a)
	}
}
