package main
import (
	"fmt"
)

//演示golang中字符串string类型的使用
func main() {
	var address string = "北京长城 110 hello world！"
	fmt.Println(address)
	//var str = "hello"
	//str[0] = 'a' //这里就不能取修改str的内容，即GO中的字符串是不可变的

	str2 := "abc\nabc"
	fmt.Println(str2)

	str3 := `
	package main
	import (
		"fmt"
		"unsafe"
	)

	//演示golang中布尔类型的使用
	func main() {
		var b = false
		fmt.Println("b=",b)
		//注意事项
		//1.bool类型占用的存储空间是1个字节
		fmt.Println("b的占用空间=", unsafe.Sizeof(b))
		//2.bool类型只能取true或者false
		c1 :=  "1"
		fmt.Println(c1)
	}
	`
	fmt.Println(str3)

	//字符串拼接方式
	var str1 = "hello" + "world"
	str1 += "haha!"
	fmt.Println(str1)
	//当一个拼接的操作很长时，怎么办，可以分行写. "+"必须放在上一行。因为go会在每行后面加一个";"
	var str5 = "hello" + "world" +
	"hello" + "world" + 
	"hello" + "world"
	fmt.Println(str5)

	var a int  // 0
	var b float32 // 0
	var c float64 // 0
	var d bool // false
	var e string // ""
	//这里的%v，表示按照变量的值输出
	fmt.Printf("a=%d,b=%v,c=%v,d=%v,e=%v\n",a,b,c,d,e)

}