package main

// 数组切片演示
import (
	"fmt"
)

func main() {
	//1，基于一个数组创建数组切片
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[5:]

	for _, v := range myArray {
		fmt.Println(v, " ")
	}

	// 修改数组元素，切片的内容也修改了
	myArray[0] = 0

	fmt.Println("----------------------")
	for _, v := range mySlice {
		fmt.Println(v, " ")
	}

	//2，直接创建
	// 2.1 初始化个数为5的数组切片，元素初始化值为0
	// mySlice1 := make([]int, 5)

	// 2.2 初始化个数为5的数组切片，元素初始化值为0，并预留10个元素的存储空间
	//mySlice1 := make([]int, 5, 10)

	// 2.3 直接创建5个元素的数组切片
	//mySlice1 := []int{1, 2, 3, 4, 5}

	// 2.4 动态增加元素
	// fmt.Println("len(mySlice1:", len(mySlice1))
	// fmt.Println("cap(mySlice1:", cap(mySlice1))

	// mySlice1 = append(mySlice1, 1, 2, 3)

	// for _, v := range mySlice1 {
	// 	fmt.Println("mySlice1----", v)
	// }

	// 2.5 基于数组切片创建新的数组切片
	/*oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	for _, v := range newSlice {
		fmt.Println("newSlice----", v)
	}*/

	// 2.6 内容复制，2个数组不一样大，会按照其中较小的那个数组切片的元素个数进行复制
	/*slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}*/

	/*copy(slice1, slice2)
	for _, v := range slice1 {
		fmt.Println("slice1 v:", v)
	}*/

	/*copy(slice2, slice1)
	for _, v := range slice2 {
		fmt.Println("slice2 v:", v)
	}*/

}
