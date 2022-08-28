package datastructures

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) ToString() string {
	var res []string
	var dummy = &ListNode{Val: 0, Next: h}
	for dummy.Next != nil {
		res = append(res, strconv.Itoa(dummy.Next.Val))
		dummy = dummy.Next
	}
	return strings.Join(res, "->")
}
