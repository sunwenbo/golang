package main
import (
	"fmt"
	_ "unsafe"  //如果我们没有使用引入的一个包，但是又不想去掉，前面加一个'_'表示忽略
	"strconv"
)

//golang中基本数据类型转成string使用
func main() {
	var num1 int = 99 
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string //空的str
	//使用第一种方式转换 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T  str=%q\n", str, str)

    str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)


	//使用第二种方式转换 strconv包函数
	var num3 int = 99 
	var num4 float64 = 23.456
	var c bool = true

	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str=%q\n", str, str)

	//说明：’f‘代表格式  10表示小数位保留位数， 64表示这个小数据是float64
	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(c)
	fmt.Printf("str type %T str=%q\n", str, str)

	//使用第三种方式转换 strconv包中有一个函数Itoa
	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str)


}