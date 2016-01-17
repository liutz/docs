package main

import "fmt"

func main() {
	// 一个中文字符在utf-8占3个字节
	str := "Hello,世界" // 依次取字符串中的字符，类型为byte
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

        // Unicode字符方式遍历
	for i, ch := range str {
		//ch的类型为rune
		fmt.Println(i, ch)
	}
}
