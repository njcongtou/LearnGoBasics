package main

import "fmt"

/**
slice 就是对array的一个view
slice 可以向后扩展
*/
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6]) // 不包含6元素的
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println(s1)

	s1 = arr[2:6]
	fmt.Println("reslice s1: ", s1)
	s2 := s1[3:5] // 取的还是arr的
	fmt.Println(s2)

	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	/**
	  slice 的操作
	  go会新开一个array，把10 放进去

	*/
	s3 := append(s2, 10)
	fmt.Println(s3)

	// 创建一个slice
	var s []int // zero value for slice is nil
	fmt.Println(s)

	for i := 0; i < 10; i++ {
		s = append(s, 2*i+1)
		fmt.Println(len(s), " ", cap(s))
	}
	fmt.Println(s)

	// 创建一个slice去view一个array
	s1 = []int{1, 2, 3, 4}
	fmt.Println(s1)

	// 删除元素
	s1 = append(s1[:2], s1[3:]...) // 4开头所有的元素，后面要加3个点
	fmt.Println(s1)

	// pop from front
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
}

func updateSlice(s []int) {
	s[0] = 100
}
