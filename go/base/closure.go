package main

// 闭包的价值
import (
	"fmt"
)

func main() {
	var j int = 5
	// 匿名函数赋值给变量a，变量a不仅表示数据，还可以表示代码
	// 变量a表示一段代码
	a := func() func() {
		// 变量i只有内部的匿名函数才可以访问，其他都访问不到，保证了i的安全性
		var i int = 10
		return func() {
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}()

	// 调用函数
	a()

	j *= 2
	a()
}
