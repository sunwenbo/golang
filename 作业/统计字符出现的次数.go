package main

import "fmt"
//qwerweirqwroiutaffab
//aaaaaaaaaaaaaaaaaaaa
func main() {
	var arr [20]byte
	var ch [26]int //用来统计字符个数

	for i :=0; i < len(arr); i++ {
		fmt.Scanf("%c",&arr[i])
	}
	//打印字母出现的次数
	for  i := 0; i < len(arr); i++ {
		ch[arr[i]-'a']++
		fmt.Println(ch)
	}
	for i := 0; i < len(ch)-1; i++ {
		if ch[i] > 0 {
			fmt.Printf("字母：%c 出现:%d次\n",'a'+i, ch[i])
		}
	}

}
