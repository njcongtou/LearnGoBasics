package main

import "fmt"

var ss = 3
var aa string = "test"
var bb = true

var (
	cc = 4
	dd = "test"
)

func main() {
	var a string
	a = "test"
	fmt.Println(a)

	var (
		b string
		c string
	)

	b = "test"
	c = "test"

	fmt.Println(b)
	fmt.Println(c)

	d := "test"
	fmt.Println(d)

	var e, f = 3, 4
	fmt.Println(e, f)

	/**
	变量类型
	1. bool, string
	2. (u)int, int8, int16, int32, int64, uintptr
	3. byte, rune: char类型
	4. float32, float64, complex64, complex128: 复合数据, 很少用，工程变量
	*/
	var test int32 = 54
	fmt.Println(test)

	/**
	强制类型转换
	*/
	var (
		c1 = 2
		c2 = "4"
	)
	c2 = string(c1)
	fmt.Println(c2)

	/**
	常量
	*/
	const filename string = "abc.txt"
	const a1, b1 = "abc.txt", "test"
	fmt.Println(a1, b1, filename)

	/**
	enum
	*/
	enums()
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
	)

	fmt.Println(cpp, java, python)

	const (
		cpp1 = iota
		java1
		python1
	)

	fmt.Println(cpp1, java1, python1)

	const (
		cpp2 = 1 << (10 * iota) // 自增表达式
		java2
		python2
	)

	fmt.Println(cpp2, java2, python2)

}
