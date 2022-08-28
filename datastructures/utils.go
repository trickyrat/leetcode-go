package datastructures

import (
	"strconv"
	"strings"
)

func CreateListNode(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	dummyHead := head
	for i := 1; i < len(nums); i++ {
		dummyHead.Next = &ListNode{Val: nums[i]}
		dummyHead = dummyHead.Next
	}
	return head
}

func CreateTreeNodeIteratively(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	sp := strings.Split(data, ",")
	n := len(sp)
	if sp[0] == "null" || n == 0 {
		return nil
	}
	val, _ := strconv.Atoi(sp[0])
	root := &TreeNode{val, nil, nil}
	var queue []*TreeNode
	queue = append(queue, root)
	index := 1
	for index < n {
		var node = queue[0]
		queue = queue[1:]
		if index > len(sp)-1 || sp[index] == "null" {
			node.Left = nil
		} else {
			leftVal, _ := strconv.Atoi(sp[index])
			leftNode := &TreeNode{leftVal, nil, nil}
			if node != nil {
				node.Left = leftNode
			}
			queue = append(queue, leftNode)
		}
		if index+1 > len(sp)-1 || sp[index+1] == "null" {
			node.Right = nil
		} else {
			rightVal, _ := strconv.Atoi(sp[index+1])
			rightNode := &TreeNode{rightVal, nil, nil}
			if node != nil {
				node.Right = rightNode
			}
			queue = append(queue, rightNode)
		}
		index += 2
	}
	return root
}
