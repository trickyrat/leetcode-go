package main

import "datastructures"

func main() {
	head := datastructures.CreateListNode([]int{1, 2, 3})
	println(head.ToString())
}
