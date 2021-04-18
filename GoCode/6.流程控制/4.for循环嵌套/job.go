package main

import "fmt"

func main() {
	chicken := 0
	count := 0
	 for cock := 0;cock <= 20; cock+=4 {
	 	for hen := 0;hen <= 33; hen++ {
	 		count++
			//代码优化 计算剩余小鸡个数 100 - 公鸡 - 母鸡
			chicken = 100 - cock - hen
			if cock * 5 + hen * 3 + chicken/3 == 100 && chicken %3 == 0 {
				fmt.Printf("公鸡:%d,母鸡:%d,小鸡:%d\n",cock,hen,chicken)
			}
	 		//for chicken := 0;chicken <= 100; chicken+=3 {
	 		//	if cock + hen + chicken == 100 && cock * 5 + hen * 3 + chicken /3 == 100{
	 		//		fmt.Printf("公鸡:%d,母鸡:%d,小鸡:%d：\n",cock,hen,chicken)
			//	}
			//}
		}
	 }
	 fmt.Println(count)
}






//func main() {
//
//	for cock := 0;cock <= 20; cock++ {
//		for hen := 0;hen <= 33; hen++ {
//			for chicken := 0;chicken <= 100; chicken+=3 {
//				if cock + hen + chicken == 100 && cock * 5 + hen * 3 + chicken /3 == 100{
//					fmt.Printf("公鸡:%d,母鸡:%d,小鸡:%d：\n",cock,hen,chicken)
//				}
//			}
//		}
//	}
//}