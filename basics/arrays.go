package main

import "fmt"

/**
  要知道数组个数
  很麻烦
  go语言一般用slice
*/
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]bool
	fmt.Println(grid)

	// print array
	for i, v := range arr3 {
		fmt.Println(i, " ", v)
	}

	printArray1(arr1)
	for i, v := range arr1 {
		fmt.Println(i, " ", v)
	}

	printArray2(&arr1)
	for i, v := range arr1 {
		fmt.Println(i, " ", v)
	}
}

// 每次都是要把数组copy一次
func printArray1(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 可以改变原来的数组里面的值了
func printArray2(arr *[5]int) {
	arr[0] = 99
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
