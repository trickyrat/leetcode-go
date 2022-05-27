package main

import (
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func createListNode(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	dummyHead := head
	for i := 1; i < len(nums); i++ {
		dummyHead.Next = &ListNode{Val: nums[i]}
		dummyHead = dummyHead.Next
	}
	return head
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
	var res = "["
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

func testArrayAndReturnInt(t *testing.T, testFunction func([]int) int, input []int, expected int) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%v) = %d; expected %d", nameof(testFunction), input, actual, expected)
	}
}

func testArrayIntAndReturnInt(t *testing.T, testFunction func(nums []int, k int) int, input []int, k, expected int) {
	actual := testFunction(input, k)
	if actual != expected {
		t.Errorf("%s(%v, %d) = %d; expected %d", nameof(testFunction), input, k, actual, expected)
	}
}

func testArrayAndReturnBool(t *testing.T, testFunction func([]int) bool, input []int, expected bool) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%v) = %t; expected %t", nameof(testFunction), input, actual, expected)
	}
}

func testArrayIntegerAndReturnArray(t *testing.T, testFunction func(array []int, target int) []int, input struct {
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

func testArrayIntegerAndStringReturnArray(t *testing.T, testFunction func(array []int, s string) []int, input []int, s string, expected []int) {
	actual := testFunction(input, s)
	for index, val := range actual {
		if val != expected[index] {
			t.Errorf("%s(%d, %s) = %v; expected %v", nameof(testFunction), input, s, actual, expected)
		}
	}
}

func testArrayAndReturnInteger(t *testing.T, testFunction func(array []string) int, input []string, expected int) {
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

func testMatrixReturnInt(t *testing.T, testFunction func([][]int) int, matrix [][]int, expected int) {
	actual := testFunction(matrix)
	if actual != expected {
		t.Errorf("%s(%v) = %d; expected %d", nameof(testFunction), matrix, actual, expected)
	}
}

func testStringAndStringArray(t *testing.T, testFunction func(string, []string) string, words string, banned []string, expected string) {
	actual := testFunction(words, banned)
	if actual != expected {
		t.Errorf("%s(%s, %v) = %s; expected %s", nameof(testFunction), words, banned, actual, expected)
	}
}

func testAddTwoNumbers(t *testing.T, num1, num2, res []int) {
	l1 := createListNode(num1)
	l2 := createListNode(num2)
	expected := createListNode(res)
	testTwoListNodeAndReturnListNode(t, addTwoNumbers, l1, l2, expected)
}

func testTwoStringsAndReturnStringFramework(t *testing.T, testFunction func(str1, str2 string) string, str1, str2, expected string) {
	actual := testFunction(str1, str2)
	if actual != expected {
		t.Errorf("%s(%s, %s) = %s; expected %s", nameof(testFunction), str1, str2, actual, expected)
	}
}

func testTwoStringArraysAndReturnStringArray(t *testing.T, testFunction func(str1, str2 []string) []string, str1, str2, expected []string) {
	actual := testFunction(str1, str2)
	for i, item := range expected {
		if item != actual[i] {
			t.Errorf("%s(%s, %s) = %s; expected %s", nameof(testFunction), str1, str2, actual, expected)
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

func testIntAndReturnString(t *testing.T, testFunction func(int) string, input int, expected string) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%d) = %s; expected %s", nameof(testFunction), input, actual, expected)
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

func testIntAndReturnIntArray(t *testing.T, testFunction func(int) []int, input int, expected []int) {
	actual := testFunction(input)
	for i, item := range actual {
		if item != expected[i] {
			t.Errorf("%s(%d) = %v; expected %v", nameof(testFunction), input, actual, expected)
		}
	}

}

func testUint32AndReturnInt(t *testing.T, testFunction func(uint32) int, input uint32, expected int) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%b) = %d; expected %d", nameof(testFunction), input, actual, expected)
	}
}

func testStringAndIntReturnString(t *testing.T, testFunction func(string, int) string, s string, numRows int, expected string) {
	actual := testFunction(s, numRows)
	if actual != expected {
		t.Errorf("%s(%s, %d) = %s; expected %s", nameof(testFunction), s, numRows, actual, expected)
	}
}

func testInputAStringAndReturnString(t *testing.T, testFunction func(string) string, input, expected string) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%s) = %s; expected %s", nameof(testFunction), input, actual, expected)
	}
}

func testStringArrayAndReturnArray(t *testing.T, testFunction func(string, [][]int) []int, input string, queries [][]int, expected []int) {
	actual := testFunction(input, queries)
	for i, item := range actual {
		if item != expected[i] {
			t.Errorf("%s(%s, %v) = %v; expected %v", nameof(testFunction), input, queries, actual, expected)
		}
	}

}

func testTwoStringAndReturnInt(t *testing.T, testFunction func(a, b string) int, a, b string, expected int) {
	actual := testFunction(a, b)
	if actual != expected {
		t.Errorf("%s(%s, %s) = %d; expected %d", nameof(testFunction), a, b, actual, expected)
	}
}

func testListNodeAndReturnListNode(t *testing.T, testFunction func(*ListNode) *ListNode, input *ListNode, expected *ListNode) {
	actual := testFunction(input)
	if actual.tostring() != expected.tostring() {
		t.Errorf("%s(%s) = %s; expected %s", nameof(testFunction), input.tostring(), actual.tostring(), expected.tostring())
	}
}

func testTwoIntegersAndReturnArray(t *testing.T, testFunction func(int, int) []int, num1, num2 int, expected []int) {
	actual := testFunction(num1, num2)
	for i, item := range actual {
		if item != expected[i] {
			t.Errorf("%s(%d, %d) = %v; expected %v", nameof(testFunction), num1, num2, actual, expected)
		}
	}
}

func testTwoIntegersAndReturnInteger(t *testing.T, testFunction func(int, int) int, num1, num2 int, expected int) {
	actual := testFunction(num1, num2)
	if actual != expected {
		t.Errorf("%s(%d, %d) = %d; expected %d", nameof(testFunction), num1, num2, actual, expected)
	}
}

func testStringAndReturnInt(t *testing.T, testFunction func(string) int, input string, expected int) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%s) = %d; expected %d", nameof(testFunction), input, actual, expected)
	}
}

func TestTwoSum(t *testing.T) {
	testArrayIntegerAndReturnArray(t, twoSum, struct {
		array  []int
		target int
	}{[]int{2, 7, 11, 15}, 9}, []int{0, 1})
	testArrayIntegerAndReturnArray(t, twoSum, struct {
		array  []int
		target int
	}{[]int{3, 2, 4}, 6}, []int{1, 2})
	testArrayIntegerAndReturnArray(t, twoSum, struct {
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
	testStringAndIntReturnString(t, zconvert, "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR")
	testStringAndIntReturnString(t, zconvert, "PAYPALISHIRING", 4, "PINALSIGYAHRPI")
	testStringAndIntReturnString(t, zconvert, "A", 1, "A")
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
	var values [][]int
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

func TestReverseListNode(t *testing.T) {
	testListNodeAndReturnListNode(t, reverseList, createListNode([]int{1, 2, 3, 4, 5}), createListNode([]int{5, 4, 3, 2, 1}))
	testListNodeAndReturnListNode(t, reverseList, createListNode([]int{1, 2}), createListNode([]int{2, 1}))
}

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	testIntAndReturnInt(t, countNumbersWithUniqueDigits, 2, 91)
	testIntAndReturnInt(t, countNumbersWithUniqueDigits, 0, 1)
}

func TestLexicalOrder(t *testing.T) {
	testIntAndReturnIntArray(t, lexicalOrder, 13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9})
	testIntAndReturnIntArray(t, lexicalOrder, 2, []int{1, 2})
}

func TestValidUtf8(t *testing.T) {
	testArrayAndReturnBool(t, validUtf8, []int{197, 130, 1}, true)
	testArrayAndReturnBool(t, validUtf8, []int{235, 140, 4}, false)
}

func TestFindSubstringWraparoundString(t *testing.T) {
	testStringAndReturnInt(t, findSubstringWraparoundString, "a", 1)
	testStringAndReturnInt(t, findSubstringWraparoundString, "cac", 2)
	testStringAndReturnInt(t, findSubstringWraparoundString, "zab", 6)
}

func TestConvertToBase7(t *testing.T) {
	testIntAndReturnString(t, convertToBase7, 100, "202")
	testIntAndReturnString(t, convertToBase7, -7, "-10")
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

func TestFindRestaurant(t *testing.T) {
	testTwoStringArraysAndReturnStringArray(t, findRestaurant, []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, []string{"Shogun"})
	testTwoStringArraysAndReturnStringArray(t, findRestaurant, []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"})
}

func TestSelfDividingNumbers(t *testing.T) {
	testTwoIntegersAndReturnArray(t, selfDividingNumbers, 1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22})
	testTwoIntegersAndReturnArray(t, selfDividingNumbers, 48, 85, []int{48, 55, 66, 77})
}

func TestUniqueMorseRepresentations(t *testing.T) {
	testArrayAndReturnInteger(t, uniqueMorseRepresentations, []string{"gin", "zen", "gig", "msg"}, 2)
	testArrayAndReturnInteger(t, uniqueMorseRepresentations, []string{"a"}, 1)
}

func TestNumberOfLines(t *testing.T) {
	testArrayIntegerAndStringReturnArray(t, numberOfLines, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz", []int{3, 60})
	testArrayIntegerAndStringReturnArray(t, numberOfLines, []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa", []int{2, 4})
}

func TestMostCommonWord(t *testing.T) {
	testStringAndStringArray(t, mostCommonWord, "Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball")
}

func TestProjectionArea(t *testing.T) {
	var grid1 [][]int
	grid1 = append(grid1, []int{1, 2})
	grid1 = append(grid1, []int{3, 4})
	var grid2 [][]int
	grid2 = append(grid2, []int{2})
	var grid3 [][]int
	grid3 = append(grid3, []int{1, 0})
	grid3 = append(grid3, []int{0, 2})
	testMatrixReturnInt(t, projectionArea, grid1, 17)
	testMatrixReturnInt(t, projectionArea, grid2, 5)
	testMatrixReturnInt(t, projectionArea, grid3, 8)
}

func TestRepeatedNTimes(t *testing.T) {
	testArrayAndReturnInt(t, repeatedNTimes, []int{1, 2, 3, 3}, 3)
	testArrayAndReturnInt(t, repeatedNTimes, []int{2, 1, 2, 5, 3, 2}, 2)
	testArrayAndReturnInt(t, repeatedNTimes, []int{5, 1, 5, 2, 5, 3, 5, 4}, 5)
}

func TestFindTheWinner(t *testing.T) {
	testTwoIntegersAndReturnInteger(t, findTheWinner, 5, 2, 3)
	testTwoIntegersAndReturnInteger(t, findTheWinner, 6, 5, 1)
}

func TestPivotIndex(t *testing.T) {
	testArrayAndReturnInt(t, pivotIndex, []int{2, 3, -1, 8, 4}, 3)
	testArrayAndReturnInt(t, pivotIndex, []int{1, -1, 4}, 2)
	testArrayAndReturnInt(t, pivotIndex, []int{2, 5}, -1)
}

func TestCountKDifference(t *testing.T) {
	testArrayIntAndReturnInt(t, countKDifference, []int{1, 2, 2, 1}, 1, 4)
	testArrayIntAndReturnInt(t, countKDifference, []int{1, 3}, 3, 0)
	testArrayIntAndReturnInt(t, countKDifference, []int{3, 2, 1, 5, 4}, 2, 3)
}

func TestMaximumDifference(t *testing.T) {
	testArrayAndReturnInt(t, maximumDifference, []int{7, 1, 5, 4}, 4)
	testArrayAndReturnInt(t, maximumDifference, []int{9, 4, 3, 2}, -1)
	testArrayAndReturnInt(t, maximumDifference, []int{1, 5, 2, 10}, 9)
}

func TestCountMaxOrSubsets(t *testing.T) {
	testArrayAndReturnInt(t, countMaxOrSubsets, []int{3, 1}, 2)
	testArrayAndReturnInt(t, countMaxOrSubsets, []int{2, 2, 2}, 7)
	testArrayAndReturnInt(t, countMaxOrSubsets, []int{3, 2, 1, 5}, 6)
}

func TestPlatesBetweenCandles(t *testing.T) {
	query1 := []int{2, 5}
	query2 := []int{5, 9}
	var queries1 [][]int
	queries1 = append(queries1, query1)
	queries1 = append(queries1, query2)
	testStringArrayAndReturnArray(t, platesBetweenCandles, "**|**|***|", queries1, []int{2, 3})

	query3 := []int{1, 17}
	query4 := []int{4, 5}
	query5 := []int{14, 17}
	query6 := []int{5, 11}
	query7 := []int{15, 16}
	var queries2 [][]int
	queries2 = append(queries2, query3)
	queries2 = append(queries2, query4)
	queries2 = append(queries2, query5)
	queries2 = append(queries2, query6)
	queries2 = append(queries2, query7)
	testStringArrayAndReturnArray(t, platesBetweenCandles, "***|**|*****|**||**|*", queries2, []int{9, 0, 0, 0, 0})
}
