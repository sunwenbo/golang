package main

import "fmt"
//if 可以嵌套 可以判断多个区间  执行效率比较低
//switch 优点执行效率低，缺点 不可以嵌套和判断多个区间
func main01() {
	// switch中的值不能为浮点数
	// 如果输入的值没有找到，默认进入default
	var w int
	fmt.Scan(&w)
	switch w {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("输入错误")
	}
}

func main() {
	var score int
	fmt.Println("请输入分数:")
	fmt.Scan(&score)
	switch score/10 {
	case 10:
		//fmt.Println("A")
		fallthrough // 让switch 执行下一个分支的代码 如果不写，就会执行到下一个分支就会自动停止
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("劝退")

	}
	i := 123
	i = i % 10
	fmt.Println(i)
}