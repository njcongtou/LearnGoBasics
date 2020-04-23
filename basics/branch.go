package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n\n", string(contents))
	}

	/**
	go 的 if else 写法
	*/
	var (
		contents2 []byte
		err2      error
	)

	if contents2, err2 = ioutil.ReadFile(filename); err != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s\n\n", string(contents2))
	}

	fmt.Printf("%s\n\n", string(contents2))

	fmt.Print(myswitch(3))
	fmt.Print(myswitch(100))
}

/**

默认有 break
除非 fallthrough
switch 后面可以不跟表达式
*/
func myswitch(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 70:
		g = "C"
	case score < 90:
		g = "B"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}

	return g
}
