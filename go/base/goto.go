package main

// 演示goto关键字使用
import (
	"fmt"
)

func myfunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func main() {
	myfunc()
}
