package main

import "fmt"

/**
  go指针简单在于不能运算



*/
func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	a, b := 3, 4
	fmt.Println(swap1(a, b))

	swap2(&a, &b)
	fmt.Println(a, b)
}

/**
1. 参数传递，值传递，还是，引用传递
    go 语言只有值传递一种方式
*/

func swap1(a, b int) (int, int) {
	return b, a
}

// 一般操作数组用slice去做
func swap2(a, b *int) {
	*a, *b = *b, *a
}
