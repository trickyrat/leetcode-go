package main

import (
	"reflect"
	"runtime"
	"strconv"
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

func testIntAndReturnBool(t *testing.T, testFunction func(int) bool, input int, expected bool) {
	actual := testFunction(input)
	if actual != expected {
		t.Errorf("%s(%d) = %t; expected %t", nameof(testFunction), input, actual, expected)
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

func TestPowerOfTwo(t *testing.T) {
	testIntAndReturnBool(t, isPowerOfTwo, 1, true)
	testIntAndReturnBool(t, isPowerOfTwo, 16, true)
	testIntAndReturnBool(t, isPowerOfTwo, 3, false)
	testIntAndReturnBool(t, isPowerOfTwo, 4, true)
	testIntAndReturnBool(t, isPowerOfTwo, 5, false)
}

func TestHammingWeight(t *testing.T) {
	testUint32AndReturnInt(t, hammingWeight, 00000000000000000000000000001011, 3)
	testUint32AndReturnInt(t, hammingWeight, 00000000000000000000000010000000, 1)
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

func TestMaximumDifference(t *testing.T) {
	testArrayAndReturnInt(t, maximumDifference, []int{7, 1, 5, 4}, 4)
	testArrayAndReturnInt(t, maximumDifference, []int{9, 4, 3, 2}, -1)
	testArrayAndReturnInt(t, maximumDifference, []int{1, 5, 2, 10}, 9)
}
