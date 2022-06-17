package main

import (
	"fmt"
	"runtime"
)

var Pi int

func init() {
	
}

func main() {
	// 数据类型
	// string、
	// bool、
	// int、uint、
	// float
	const (
		// 可用于枚举、iota 是一个预先声明的标识符，表示（通常用括号括起来的）const 声明中当前 const 规范的无类型整数序数。它是零索引的。
		// 假如 a, b = iota + 1, iota + 2,
		// 像这样赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
		a = 1 << iota * 3
		b
		c
		d
		e
		f
	)

	fmt.Print(a, b, c, d, e, f, "\n")
	// const 属于声名常量
	const age = 1
	const nameC string = "1"
	const nameA, nameB = "张三", "李四"
	// var 属于声明变量
	var var1 = 32
	var var2 uint8 = 31
	var3 := 33
	fmt.Print(var1, var2, var3, "\n")
	// 数组
	var arr = [3]int{123, 12, 12}
	arr1 := [3]int{123, 12, 12}
	fmt.Print(arr, "-", arr1, "\n")
	fmt.Print("rain\n")
	goos := runtime.GOOS
	fmt.Print(goos + "\n")
}

func sum(a, b int) int {
	return a + b
}

func bianLiang(a, b int) int {
	var name = 123
	name = a
	return name
}
