package leetcode

import (
	"datastructures"
	"strconv"
	"strings"
)

const mask1, mask2 = 1 << 7, 1<<7 | 1<<6

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

func calculateDepth(root *datastructures.TreeNode) int {
	height := -1
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		height++
		temp := queue
		queue = nil
		for _, node := range temp {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return height
}

func round(value float64, digits int) float64 {
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', digits, 64), 64)
	return value
}
