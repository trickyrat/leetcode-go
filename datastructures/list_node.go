package datastructures

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) toString() string {
	var res = "["
	res += strconv.Itoa(h.Val)
	for h.Next != nil {
		res += strconv.Itoa(h.Val)
		h = h.Next
	}
	res += "]"
	return res
}
