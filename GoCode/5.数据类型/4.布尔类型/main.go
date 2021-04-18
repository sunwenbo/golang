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