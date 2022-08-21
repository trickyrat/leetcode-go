package leetcode

import (
	"datastructures"
	"strconv"
	"strings"
)

const mask1, mask2 = 1 << 7, 1<<7 | 1<<6

func createListNode(nums []int) *datastructures.ListNode {
	head := &datastructures.ListNode{Val: nums[0]}
	dummyHead := head
	for i := 1; i < len(nums); i++ {
		dummyHead.Next = &datastructures.ListNode{Val: nums[i]}
		dummyHead = dummyHead.Next
	}
	return head
}

func createTreeNodeWithBFS(data string) *datastructures.TreeNode {
	if len(data) == 0 {
		return nil
	}
	sp := strings.Split(data, ",")
	n := len(sp)
	if sp[0] == "null" || n == 0 {
		return nil
	}
	val, _ := strconv.Atoi(sp[0])
	root := &datastructures.TreeNode{val, nil, nil}
	var queue []*datastructures.TreeNode
	queue = append(queue, root)
	index := 1
	for index < n {
		var node = queue[0]
		queue = queue[1:]
		if index > len(sp)-1 || sp[index] == "null" {
			node.Left = nil
		} else {
			leftVal, _ := strconv.Atoi(sp[index])
			leftNode := &datastructures.TreeNode{leftVal, nil, nil}
			if node != nil {
				node.Left = leftNode
			}
			queue = append(queue, leftNode)
		}
		if index+1 > len(sp)-1 || sp[index+1] == "null" {
			node.Right = nil
		} else {
			rightVal, _ := strconv.Atoi(sp[index+1])
			rightNode := &datastructures.TreeNode{rightVal, nil, nil}
			if node != nil {
				node.Right = rightNode
			}
			queue = append(queue, rightNode)
		}
		index += 2
	}
	return root
}

func createTreeNodeWithDFS(data string) *datastructures.TreeNode {
	sp := strings.Split(data, ",")
	var build func() *datastructures.TreeNode
	build = func() *datastructures.TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &datastructures.TreeNode{val, build(), build()}
	}
	return build()
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func isSelfDividing(num int) bool {
	for x := num; x > 0; x /= 10 {
		if d := x % 10; d == 0 || num%d != 0 {
			return false
		}
	}
	return true
}

func parseComplexNumber(num string) (real, imag int) {
	i := strings.IndexByte(num, '+')
	real, _ = strconv.Atoi(num[:i])
	imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
	return
}

func getBytes(num int) int {
	if num&mask1 == 0 {
		return 1
	}
	n := 0
	for mask := mask1; num&mask != 0; mask >>= 1 {
		n++
		if n > 4 {
			return -1
		}
	}
	if n >= 2 {
		return n
	}
	return -1
}
