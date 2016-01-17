package main

import (
	"fmt"
)

// 数组是一个值类型，在参数传递时产生一次复制动作
func modify(array [5]int) {
	array[0] = 10
	fmt.Println("in modify(),array values:", array)
}

func main() {
	// 定义5个元素的数组
	// var array [5]int = [5]int{1, 2, 5, 4, 3}	
	array := [5]int{1, 2, 5, 4, 3}
	//arrLength := len(array)
	// for i := 0; i < arrLength; i++ {
	// 	fmt.Println("Element", i, "of array is", array[i])
	// }

	modify(array)

	for i, v := range array {
		fmt.Println("Array element[", i, "]=", v)
	}

}
