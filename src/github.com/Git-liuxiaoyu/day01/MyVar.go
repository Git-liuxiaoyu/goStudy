// 变量
package main

import "fmt"

// 变量块，必须指定 类型
// 也可以直接赋值 _name = "rain"
var (
	_name string
)

// 主函数
func main() {
	// 定义变量的三种写法
	var _var1 string = "_var1"
	var _var2 = "_var2"
	// 类型推导
	_var3 := "_var3"
	// 定义变量后必须给用到它们，不然会报错，但是变量块里面的就不一样
	fmt.Println(_var1)
	fmt.Println(_var2)
	fmt.Println(_var3)
}
