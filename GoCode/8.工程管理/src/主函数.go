package main

import "userinfo"
//如果调用不通级别目录，需要导入包


/*
src目录：用于以代码包的形式组织并保存Go源码文件。比如go .c .h .s等
pkg目录：用于存放经由go install命令构建安装后的代码包（包含Go库源码文件）的"a"归档文件。
bin目录：与pkg类似，在通过go install 命令完成安装后，保存由Go命令源码文件生成的可执行文件
 */
//全局变量
var abc int = 123

func main() {
	add(10,20)
	//userinfo.Login()
	userinfo.Login()
	userinfo.DeleteUser()
}