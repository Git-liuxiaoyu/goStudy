package main

import "fmt"

// 函数外面只能放标识符的声明

// 数据类型
// string、
// bool、
// int、uint、
// float

var (
	_name string
)

const (
	_var1 = iota
	_var2
	_var3
)

// 程序的入口
func main() {
	_name = "123"
	fmt.Println("name:" + _name)
	fmt.Print(_var1, _var2, _var3)
}
