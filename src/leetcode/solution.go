package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.Two Sum
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

// 2. Add Two Numbers
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

// 191. Number of 1 Bits
func hammingWeight(num uint32) int {
	var res int = 0
	for {
		if num == 0 {
			break
		}
		num &= num - 1
		res++
	}
	return res
}

//231. Power of Two
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
