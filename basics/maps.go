package main

import "fmt"

/**
  map 不保证key的顺序
  除了slice, map, function的内建类型都可以作为key
  struct 类型不包含上面的字段，也可以作为key
*/
func main() {

	m := map[string]string{
		"name": "congwang",
		"age":  "21",
	}

	m2 := make(map[string]int)

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m["name"])

	// 判断key在不在
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("cause key does not exist.")
	}

	// 删除元素
	delete(m, "name")
	fmt.Println(m)

}
