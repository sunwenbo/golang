package main
import (
	"fmt"
	_ "unsafe"  //如果我们没有使用引入的一个包，但是又不想去掉，前面加一个'_'表示忽略
	"strconv"
)

//golang中基本string转换基础类型
func main() {
	//字符转布尔值
	var str string = "true"
	var b bool
	//	b , _ = strconv.ParseBool(str)
	//说明：
	//1.这个函数会返回两个值，(value bool, err error)
	//2.因为我只想获取到value bool,不想获取err 所有使用_忽略掉
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n",b ,b)

	//字符转布尔值
	var str2 string = "1234567890"
	var n1 int64
	n1, _ = strconv.ParseInt(str2,10,64)  //10进制，64代表int64
	fmt.Printf("n1 type %T n1=%v\n",n1 ,n1)

	//字符转浮点数
	var str3 string = "123.456"
	var f1 float64
	var f2 float32
	f1, _ = strconv.ParseFloat(str3, 64)
	//因为返回的是int64或者是float64，如希望要得到int32，float32等需要再次进行转换
	f2 = float32(f1)
	fmt.Printf("f1 type %T f1=%v\n",f1 ,f1)
	fmt.Printf("f2 type %T f2=%v\n",f2 ,f2)
}