package main

import (
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func createListNode(nums []int) ListNode {
	head := &ListNode{Val: nums[0]}
	dummyHead := head
	for i := 1; i < len(nums); i++ {
		dummyHead.Next = &ListNode{Val: nums[i]}
		dummyHead = dummyHead.Next
	}
	return *head
}

func createTreeNodeWithBFS(data string) *TreeNode {
	sp := strings.Split(data, ",")
	if sp[0] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(sp[0])
	root := &TreeNode{val, nil, nil}
	var queue []*TreeNode
	queue = append(queue, root)
	index := 1
	for index < len(sp) {
		var node = queue[0]
		queue = queue[1:]
		leftStr := sp[index]
		rightStr := sp[index+1]
		if leftStr != "null" {
			leftVal, _ := strconv.Atoi(leftStr)
			leftNode := &TreeNode{leftVal, nil, nil}
			if node != nil {
				node.Left = leftNode
			}
			queue = append(queue, leftNode)
		}
		if rightStr != "null" {
			rightVal, _ := strconv.Atoi(rightStr)
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

func createTreeNodeWithDFS(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}

func (h *ListNode) tostring() string {
	var res string = "["
	res += strconv.Itoa(h.Val)
	for h.Next != nil {
		res += strconv.Itoa(h.Val)
		h = h.Next
	}
	res += "]"
	return res
}

func nameof(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func testArrayAndReturnString(t *testing.T, testFunction func([]int) string, input []int, expected string) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%v) = %s; expected %s", nameof(testFunction), input, actual, expected)
	}
}

func testInputArrayIntAndReturnIntFramework(t *testing.T, testFunction func(nums []int, k int) int, input []int, k, expected int) {
	actual := testFunction(input, k)
	if actual != expected {
		t.Errorf("%s(%v, %d) = %d; expected %d", nameof(testFunction), input, k, actual, expected)
	}
}

func testArrayAndReturnInt(t *testing.T, testFunction func([]int) int, input []int, expected int) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%v) = %d; expected %d", nameof(testFunction), input, actual, expected)
	}
}

func testTwoListNodeAndReturnListNode(t *testing.T, testFunction func(*ListNode, *ListNode) *ListNode, l1, l2, expected *ListNode) {
	actual := testFunction(l1, l2)
	actualStr := actual.tostring()
	expectedStr := expected.tostring()
	if actualStr != expectedStr {
		t.Errorf("%s(%s, %s) = %s; expected %s", nameof(testFunction), l1.tostring(), l2.tostring(), actualStr, expectedStr)
	}
}

func testTwoStringsAndReturnStringFramework(t *testing.T, testFunction func(str1, str2 string) string, str1, str2, expected string) {
	actual := testFunction(str1, str2)
	if actual != expected {
		t.Errorf("%s(%s, %s) = %s; expected %s", nameof(testFunction), str1, str2, actual, expected)
	}
}

func testArrayIntegerAndReturnArrayFramework(t *testing.T, testFunction func(array []int, target int) []int, input struct {
	array  []int
	target int
}, expected []int) {
	actual := testFunction(input.array, input.target)
	for index, val := range actual {
		if val != expected[index] {
			t.Errorf("%s(%d, %d) = %v; expected %v", nameof(testFunction), input.array, input.target, actual, expected)
		}
	}
}

func testTreeNodeAndReturnArrayOfArray(t *testing.T, testFunction func(root *TreeNode, targetNum int) [][]int, input *TreeNode, targetNum int, expected [][]int) {
	actual := testFunction(input, targetNum)
	for i, item := range expected {
		for j, elem := range actual[i] {
			if elem != item[j] {
				t.Errorf("%s(%v, %d) = %d; expected %d", nameof(testFunction), input, targetNum, elem, item[j])
			}
		}
	}

}

func testIntAndReturnBool(t *testing.T, testFunction func(int) bool, input int, expected bool) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%d) = %t; expected %t", nameof(testFunction), input, actual, expected)
	}
}

func testIntAndReturnInt(t *testing.T, testFunction func(int) int, input, expected int) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%d) = %d; expected %d", nameof(testFunction), input, actual, expected)
	}
}

func testUint32AndReturnInt(t *testing.T, testFunction func(uint32) int, input uint32, expected int) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%b) = %d; expected %d", nameof(testFunction), input, actual, expected)
	}
}

func testInputStringAndIntReturnString(t *testing.T, testFunction func(string, int) string, s string, numRows int, expected string) {
	actual := testFunction(s, numRows)
	if actual != expected {
		t.Errorf("%s(%s, %d) = %s; expected %s", nameof(testFunction), s, numRows, actual, expected)
	}
}

func testAddTwoNumbers(t *testing.T, num1, num2, res []int) {
	l1 := createListNode(num1)
	l2 := createListNode(num2)
	expected := createListNode(res)
	testTwoListNodeAndReturnListNode(t, addTwoNumbers, &l1, &l2, &expected)
}

func testInputAStringAndReturnString(t *testing.T, testFunction func(string) string, input, expected string) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%s) = %s; expected %s", nameof(testFunction), input, actual, expected)
	}
}

func testTwoStringAndReturnInt(t *testing.T, testFunction func(a, b string) int, a, b string, expected int) {
	actual := testFunction(a, b)
	if actual != expected {
		t.Errorf("%s(%s, %s) = %d; expected %d", nameof(testFunction), a, b, actual, expected)
	}
}

func TestTwoSum(t *testing.T) {
	testArrayIntegerAndReturnArrayFramework(t, twoSum, struct {
		array  []int
		target int
	}{[]int{2, 7, 11, 15}, 9}, []int{0, 1})
	testArrayIntegerAndReturnArrayFramework(t, twoSum, struct {
		array  []int
		target int
	}{[]int{3, 2, 4}, 6}, []int{1, 2})
	testArrayIntegerAndReturnArrayFramework(t, twoSum, struct {
		array  []int
		target int
	}{[]int{3, 3}, 6}, []int{0, 1})
}

func TestAddTwoNumbers(t *testing.T) {
	testAddTwoNumbers(t, []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8})
	testAddTwoNumbers(t, []int{0}, []int{0}, []int{0})
	testAddTwoNumbers(t, []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1})
}

func TestZConvert(t *testing.T) {
	testInputStringAndIntReturnString(t, zconvert, "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR")
	testInputStringAndIntReturnString(t, zconvert, "PAYPALISHIRING", 4, "PINALSIGYAHRPI")
	testInputStringAndIntReturnString(t, zconvert, "A", 1, "A")
}

func TestReverseInt(t *testing.T) {
	testIntAndReturnInt(t, reverseInt, 123, 321)
	testIntAndReturnInt(t, reverseInt, -123, -321)
	testIntAndReturnInt(t, reverseInt, 120, 21)
	testIntAndReturnInt(t, reverseInt, 100, 1)
	testIntAndReturnInt(t, reverseInt, 0, 0)
	testIntAndReturnInt(t, reverseInt, -2147483649, 0)
	testIntAndReturnInt(t, reverseInt, -2147483648, 0)
	testIntAndReturnInt(t, reverseInt, -2147483647, 0)
	testIntAndReturnInt(t, reverseInt, 2147483647, 0)
	testIntAndReturnInt(t, reverseInt, 2147483648, 0)
	testIntAndReturnInt(t, reverseInt, 2147483649, 0)
}

func TestPathSum(t *testing.T) {
	row1 := []int{5, 4, 11, 2}
	row2 := []int{5, 8, 4, 5}
	values := [][]int{}
	values = append(values, row1)
	values = append(values, row2)

	testTreeNodeAndReturnArrayOfArray(t, pathSum, createTreeNodeWithBFS("5,4,8,11,null,13,4,7,2,null,null,5,1"), 22, values)
	testTreeNodeAndReturnArrayOfArray(t, pathSum, createTreeNodeWithBFS("1,2,3"), 5, [][]int{})
}

func TestHammingWeight(t *testing.T) {
	testUint32AndReturnInt(t, hammingWeight, 00000000000000000000000000001011, 3)
	testUint32AndReturnInt(t, hammingWeight, 00000000000000000000000010000000, 1)
}

func TestPowerOfTwo(t *testing.T) {
	testIntAndReturnBool(t, isPowerOfTwo, 1, true)
	testIntAndReturnBool(t, isPowerOfTwo, 16, true)
	testIntAndReturnBool(t, isPowerOfTwo, 3, false)
	testIntAndReturnBool(t, isPowerOfTwo, 4, true)
	testIntAndReturnBool(t, isPowerOfTwo, 5, false)
}

func TestFindLUSLength(t *testing.T) {
	testTwoStringAndReturnInt(t, findLUSLength, "aba", "cdc", 3)
	testTwoStringAndReturnInt(t, findLUSLength, "aaa", "bbb", 3)
	testTwoStringAndReturnInt(t, findLUSLength, "aaa", "aaa", -1)
}

func TestComplexNumberMultiply(t *testing.T) {
	testTwoStringsAndReturnStringFramework(t, complexNumberMultiply, "1+1i", "1+1i", "0+2i")
	testTwoStringsAndReturnStringFramework(t, complexNumberMultiply, "1+-1i", "1+-1i", "0+-2i")
}

func TestReverseOnlyLetters(t *testing.T) {
	testInputAStringAndReturnString(t, reverseOnlyLetters, "ab-cd", "dc-ba")
	testInputAStringAndReturnString(t, reverseOnlyLetters, "a-bC-dEf-ghIj", "j-Ih-gfE-dCba")
	testInputAStringAndReturnString(t, reverseOnlyLetters, "Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!")
}

func TestOptimalDivision(t *testing.T) {
	testArrayAndReturnString(t, optimalDivision, []int{1000, 100, 10, 2}, "1000/(100/10/2)")
}

func TestCountKDifference(t *testing.T) {
	testInputArrayIntAndReturnIntFramework(t, countKDifference, []int{1, 2, 2, 1}, 1, 4)
	testInputArrayIntAndReturnIntFramework(t, countKDifference, []int{1, 3}, 3, 0)
	testInputArrayIntAndReturnIntFramework(t, countKDifference, []int{3, 2, 1, 5, 4}, 2, 3)
}

func TestMaximumDifference(t *testing.T) {
	testArrayAndReturnInt(t, maximumDifference, []int{7, 1, 5, 4}, 4)
	testArrayAndReturnInt(t, maximumDifference, []int{9, 4, 3, 2}, -1)
	testArrayAndReturnInt(t, maximumDifference, []int{1, 5, 2, 10}, 9)
}
