package main

import (
	"errors"
	"fmt"
)

// 1.演示函数返回多值
func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}

	return a + b, nil
}

// 2.任意类型的可变参数
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		// 3.判断变量是什么类型
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is a int64 value.")
		default:
			fmt.Println(arg, "is an unkown type.")
		}
	}
}

func main() {

	// 4.返回错误，该怎么处理
	c, err := Add(10, 1)

	if err != nil {
		fmt.Println("please input non-negative numbers")
	} else {
		fmt.Println(c)
	}

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	// 5.定义匿名函数并赋值给变量f
	f := func(x, y int) int {
		return x + y
	}
	// 打印函数指针
	fmt.Println(f)

	f1 := func(x, y int) int {
		return x * y
	}(2, 5)

	fmt.Println(f1)

}
