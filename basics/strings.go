package main

import (
	"fmt"
	"unicode/utf8"
)

/**
  rune 相当于 go 的char



*/
func main() {
	s := "Yes我爱慕课网!" // UTF-8 encoding
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

}
