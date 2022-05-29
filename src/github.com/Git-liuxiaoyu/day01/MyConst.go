// 常量
package main

import "fmt"

// 批量声明
const (
	statusOK = 200
	notFound = 400
)

// 第二种写法
// 这一种的如果第一个赋值了，下边的没有赋值，那么var2、var3默认就是和上边var1一样
const (
	var1 = "1"
	var2
	var3
)

// iota go语言的计数器，默认值为0,
const (
	var4 = iota
	var5
	var6
	_ // 匿名变量 忽然iota,但是也夺走了iota的索引3，var7后面没有跟值，所以还是会用iota的索引 var7 = 4
	var7
)

// 当const常量块中出现iota时，默认值为0，之后每新增一列iota都会+1，但是如果iota不在第一行的话，那么默认值iota为出现在所属行的下标
const (
	a0 = 100
	a1        // 100
	a2        // 100
	a3 = iota // 3
	a4        // 4
	a5        // 5
	a6        // 6
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// 定义常量的两种写法
	const name string = "rain"
	const age = 20
	fmt.Println(var1, var2, var3)
	fmt.Println(var4, var5, var6, var7)
	fmt.Println(a0, a1, a2, a3, a4, a5, a6)
	fmt.Println(KB, MB, GB, TB, PB)
}
