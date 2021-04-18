/*
1.gopath目录
GoCode		里面每一个子目录，就是一个包。包内是Go的源码文件
pkg		编译后生成的，包的目标文件
bin		生成的可执行文件

2.import
	import "fmt"告诉Go编译器这个程序需要使用fmt包的函数，fmt包实现了格式化IO的函数可以是相对路径也可以是绝对路径，推荐使用绝对路径（起始于工程根目录）
	1.点操作
	2.别名操作
	3._操作

(1)go 文件的后缀是.go
(2)package main
   表示该hello.go 文件所在的包是main，在go中，每个文件必须属于一个包。
(3)import fmt
   表示：引入一个包，报名fmt，引入该包后，就可以使用fmt包的函数，比如：fmt.Print
(4) func main() {

}
	func 是一个关键字，表示一个函数
	main是函数名，是一个主函数，即我们程序入口
(5)fmt.Print("hello world!")


运行一个go程序    go run first-go.go
编译运行 go build first-go.go  
执行二进制可执行文件 ./first-go 生成的二进制文件较大 #实际生产环境先编译后执行二进制可执行文件

编译和运行的说明，我们可以指定生成的二进制可执行文件名称
go build -o myhello first-go.go 

Go开发注意事项：
1.Go源文件以“go”为扩展名
2.Go应用程序的执行入口是main()函数。
3.Go语言严格区分大小写
4.Go方法由一条条语句构成，每个语句后不需要分号（Go语言会在每行后自动加分号），这也是体现出Golang的简洁性
5.Go编译器是一行行进行编译的，因为我们一行就写一条语句，不能把多行语句写在同一行否则会报错。
6.Go语言定义的变量或者import的包如果没有使用到，代码不能编译通过。
7.打括号都是成对出现对，缺一不可。
*/

package main
import "fmt"

func main() {
	//var num = 10  定义后对变量必须使用，否则会报错
	fmt.Println("GO,GO,GO!!!")
	fmt.Println("Hello, World!")
}
