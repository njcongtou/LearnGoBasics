package main

import (
	"fmt"
	"reflect"
	"runtime"
)

/**
函数多返回值就是用来返回error的

*/
func main() {
	fmt.Println(div1(10, 3))

	fmt.Println(div1(11, 3))

	// 匿名函数
	// go 没有花哨的lambda表达式，这样匿名函数就是lambda了
	fmt.Println(apply(func(a, b int) int {
		return a + b
	}, 12, 4))
}

/**
ex 1. 这个就是标准的写法
*/
func div1(a, b int) (q, r int) {
	return a / b, a % b
}

/**
ex 2.
因为在返回值里面已经定义了
所以只要在函数体里面赋值，就不用在return里面写了
*/
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func test(a, b int) (int, error) {
	var e = 3
	var f = 5
	return e * f, nil
}

/**
ex 3.
用函数式编程重写下面的func
*/
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	// 用反射去拿下函数的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Println("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

/*
ex 4.
go 语言只有可变参数列表，没有其他的花哨东西
*/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
