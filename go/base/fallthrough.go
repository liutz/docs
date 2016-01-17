package main

import (
	"fmt"
)

// fallthrough关键字，才会继续执行下一个case
func main() {
	i := 3
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4, 5, 6:
		fmt.Println("4,5,6")
	default:
		fmt.Println("default")
	}
}
