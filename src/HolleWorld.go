package main

import "fmt"

func main() {
	// 数据类型
	// string、
	// bool、
	// int、uint、
	// float

	// const 属于声名常量
	const age int = 1
	const boolA bool = true
	const nameC string = "1"
	const nameA, nameB = "张三", "李四"
	// var 属于声明变量
	var age_1 uint8 = 31
	var age_2 = 32
	age_3 := 33
	var age_4, age_5, age_6 int = 31, 32, 33
	fmt.Print(age_1, age_2, age_3)
	fmt.Print(age_4, age_5, age_6)
	fmt.Printf("name%v\n", age_2)
	// go语言的占位符

	// 数组
	arr := [6]int{123, 12, 12}
	modifyArr(&arr)
	modifyArr(&arr)
	fmt.Print(arr)
	fmt.Print(arr)
}

func modifyArr(a *[6]int) {
	a[1] = 20
}
