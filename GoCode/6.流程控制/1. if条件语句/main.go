package main

import "fmt"

func main01() {
	var score int
	fmt.Scan(&score)
	// if条件判断根据是否满足条件指向对应的代码

	if score > 700 {
		fmt.Println("我要上清华")
	} else {
		fmt.Println("我要上蓝翔")
	}
}

func main02() {
	var score int
	fmt.Scan(&score)
	if score >= 700 {
		fmt.Println("我要上清华")
		if score > 720 {
			fmt.Println("我要学习挖掘机")
		} else if score > 710 {
			fmt.Println("我要学习美容美发")
		} else {
			fmt.Println("我要学习计算机")
		}
	} else if score >= 680 {
		fmt.Println("我要上北大")
		if score > 690 {
			fmt.Println("我要学习考古")
		} else if score > 685 {
			fmt.Println("我要学习占卜")
		}
	} else if score >= 650 {
		fmt.Println("我要上交大")
	} else {
		fmt.Println("我要上蓝翔")
	}
}

func main03() {
	a := 100
	b := 200
	c := 90
	if a > b {
		if a > c {
			fmt.Println("a重")
		} else {
			fmt.Println("c重")
		}
	} else {
		if b > c {
			fmt.Println("b重")
		} else {
			fmt.Println("c重")
		}
	}
}

func main() {
	a1 := 100
	a2 := 300
	a3 := 200
	if a1 > a2 {
		if a1 > a3 {
			println("a1重",a1)
		} else {
			println("a3重",a3)
		}
	} else {
		if a2 > a3 {
			println("a2",a2)
		} else {
			println("a3",a3)
		}
	}
}
















