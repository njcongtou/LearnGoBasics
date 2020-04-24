package main

import "fmt"

/**
  只有封装
  不支持继承和多态
  go 没有class 只有struct
  go 没有构造函数
  有时候想要控制构造，就加上工厂函数
  不用考虑对象是在哪里分配
*/

type treeNode struct {
	value       int
	left, right *treeNode
}

/* 相当于java里面的this, receiver */
/*
    值接受者是go语言特有
   用法：调用的时候会copy一个treeNode
        不能改变里面内容
*/
func (node treeNode) print() {
	fmt.Println(node.value)
}

/*
    指针接受者
   用法：调用的时候不会copy一个treeNode
        能改变里面内容
        结构过大也用这个
        一致性，全用一种
*/
func (node *treeNode) print2(value int) {
	fmt.Println(node.value)
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	/**
	  1. 结构体的创建
	*/
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.print()

}
