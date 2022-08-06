package main

import (
	"github.com/stretchr/testify/assert"
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

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}

func TestAddTwoNumbers(t *testing.T) {
	assert.Equal(t, addTwoNumbers(createListNode([]int{2, 4, 3}), createListNode([]int{5, 6, 4})), createListNode([]int{7, 0, 8}))
	assert.Equal(t, addTwoNumbers(createListNode([]int{0}), createListNode([]int{0})), createListNode([]int{0}))
	assert.Equal(t, addTwoNumbers(createListNode([]int{9, 9, 9, 9, 9, 9, 9}), createListNode([]int{9, 9, 9, 9})), createListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}))
}

func TestZConvert(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", zconvert("PAYPALISHIRING", 3))
	assert.Equal(t, "PINALSIGYAHRPI", zconvert("PAYPALISHIRING", 4))
	assert.Equal(t, "A", zconvert("A", 1))
}

func TestReverseInt(t *testing.T) {
	assert.Equal(t, 321, reverseInt(123))
	assert.Equal(t, -321, reverseInt(-123))
	assert.Equal(t, 21, reverseInt(120))
	assert.Equal(t, 1, reverseInt(100))
	assert.Equal(t, 0, reverseInt(0))
	assert.Equal(t, 0, reverseInt(-2147483649))
	assert.Equal(t, 0, reverseInt(-2147483648))
	assert.Equal(t, 0, reverseInt(-2147483647))
	assert.Equal(t, 0, reverseInt(2147483647))
	assert.Equal(t, 0, reverseInt(2147483648))
	assert.Equal(t, 0, reverseInt(2147483649))
}

func TestPathSum(t *testing.T) {
	row1 := []int{5, 4, 11, 2}
	row2 := []int{5, 8, 4, 5}
	var values [][]int
	values = append(values, row1)
	values = append(values, row2)
	assert.Equal(t, values, pathSum(createTreeNodeWithBFS("5,4,8,11,null,13,4,7,2,null,null,5,1"), 22))
	assert.Equal(t, [][]int(nil), pathSum(createTreeNodeWithBFS("1,2,3"), 5))
}

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(00000000000000000000000000001011))
	assert.Equal(t, 1, hammingWeight(00000000000000000000000010000000))
}

func TestPowerOfTwo(t *testing.T) {
	assert.True(t, isPowerOfTwo(1))
	assert.True(t, isPowerOfTwo(16))
	assert.False(t, isPowerOfTwo(3))
	assert.True(t, isPowerOfTwo(4))
	assert.False(t, isPowerOfTwo(5))
}

func TestReverseListNode(t *testing.T) {
	assert.Equal(t, createListNode([]int{5, 4, 3, 2, 1}), reverseList(createListNode([]int{1, 2, 3, 4, 5})))
	assert.Equal(t, createListNode([]int{2, 1}), reverseList(createListNode([]int{1, 2})))
}

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	assert.Equal(t, 91, countNumbersWithUniqueDigits(2))
	assert.Equal(t, 1, countNumbersWithUniqueDigits(0))
}

func TestLexicalOrder(t *testing.T) {
	assert.Equal(t, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}, lexicalOrder(13))
	assert.Equal(t, []int{1, 2}, lexicalOrder(2))
}

func TestValidUtf8(t *testing.T) {
	assert.True(t, validUtf8([]int{197, 130, 1}))
	assert.False(t, validUtf8([]int{235, 140, 4}))
}

func TestFindSubstringWraparoundString(t *testing.T) {
	assert.Equal(t, 1, findSubstringWraparoundString("a"))
	assert.Equal(t, 2, findSubstringWraparoundString("cac"))
	assert.Equal(t, 6, findSubstringWraparoundString("zab"))
}

func TestConvertToBase7(t *testing.T) {
	assert.Equal(t, "202", convertToBase7(100))
	assert.Equal(t, "-10", convertToBase7(-7))
}

func TestFindLUSLength(t *testing.T) {
	assert.Equal(t, 3, findLUSLength("aba", "cdc"))
	assert.Equal(t, 3, findLUSLength("aaa", "bbb"))
	assert.Equal(t, -1, findLUSLength("aaa", "aaa"))
}

func TestComplexNumberMultiply(t *testing.T) {
	assert.Equal(t, "0+2i", complexNumberMultiply("1+1i", "1+1i"))
	assert.Equal(t, "0+-2i", complexNumberMultiply("1+-1i", "1+-1i"))
}

func TestReverseOnlyLetters(t *testing.T) {
	assert.Equal(t, "dc-ba", reverseOnlyLetters("ab-cd"))
	assert.Equal(t, "j-Ih-gfE-dCba", reverseOnlyLetters("a-bC-dEf-ghIj"))
	assert.Equal(t, "Qedo1ct-eeLg=ntse-T!", reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}

func TestOptimalDivision(t *testing.T) {
	assert.Equal(t, "1000/(100/10/2)", optimalDivision([]int{1000, 100, 10, 2}))
}

func TestFindRestaurant(t *testing.T) {
	assert.Equal(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	assert.Equal(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
}

func TestSelfDividingNumbers(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}, selfDividingNumbers(1, 22))
	assert.Equal(t, []int{48, 55, 66, 77}, selfDividingNumbers(48, 85))
}

func TestUniqueMorseRepresentations(t *testing.T) {
	assert.Equal(t, 2, uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	assert.Equal(t, 1, uniqueMorseRepresentations([]string{"a"}))
}

func TestNumberOfLines(t *testing.T) {
	assert.Equal(t, []int{3, 60}, numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	assert.Equal(t, []int{2, 4}, numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func TestMostCommonWord(t *testing.T) {
	assert.Equal(t, "ball", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
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
	assert.Equal(t, 17, projectionArea(grid1))
	assert.Equal(t, 5, projectionArea(grid2))
	assert.Equal(t, 8, projectionArea(grid3))
}

func TestRepeatedNTimes(t *testing.T) {
	assert.Equal(t, 3, repeatedNTimes([]int{1, 2, 3, 3}))
	assert.Equal(t, 2, repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	assert.Equal(t, 5, repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}

func TestMinSubsequence(t *testing.T) {
	assert.Equal(t, []int{10, 9}, minSubsequence([]int{4, 3, 10, 9, 8}))
	assert.Equal(t, []int{7, 7, 6}, minSubsequence([]int{4, 4, 7, 6, 7}))
	assert.Equal(t, []int{6}, minSubsequence([]int{6}))
}

func TestStringMatching(t *testing.T) {
	assert.Equal(t, []string{"as", "hero"}, stringMatching([]string{"mass", "as", "hero", "superhero"}))
	assert.Equal(t, []string{"et", "code"}, stringMatching([]string{"leetcode", "et", "code"}))
	assert.Equal(t, []string(nil), stringMatching([]string{"blue", "green", "bu"}))
}

func TestFindTheWinner(t *testing.T) {
	assert.Equal(t, 3, findTheWinner(5, 2))
	assert.Equal(t, 1, findTheWinner(6, 5))
}

func TestPivotIndex(t *testing.T) {
	assert.Equal(t, 3, pivotIndex([]int{2, 3, -1, 8, 4}))
	assert.Equal(t, 2, pivotIndex([]int{1, -1, 4}))
	assert.Equal(t, -1, pivotIndex([]int{2, 5}))
}

func TestCountKDifference(t *testing.T) {
	assert.Equal(t, 4, countKDifference([]int{1, 2, 2, 1}, 1))
	assert.Equal(t, 0, countKDifference([]int{1, 3}, 3))
	assert.Equal(t, 3, countKDifference([]int{3, 2, 1, 5, 4}, 2))
}

func TestMaximumDifference(t *testing.T) {
	assert.Equal(t, 4, maximumDifference([]int{7, 1, 5, 4}))
	assert.Equal(t, -1, maximumDifference([]int{9, 4, 3, 2}))
	assert.Equal(t, 9, maximumDifference([]int{1, 5, 2, 10}))
}

func TestCountMaxOrSubsets(t *testing.T) {
	assert.Equal(t, 2, countMaxOrSubsets([]int{3, 1}))
	assert.Equal(t, 7, countMaxOrSubsets([]int{2, 2, 2}))
	assert.Equal(t, 6, countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func TestPlatesBetweenCandles(t *testing.T) {
	query1 := []int{2, 5}
	query2 := []int{5, 9}
	var queries1 [][]int
	queries1 = append(queries1, query1)
	queries1 = append(queries1, query2)
	assert.Equal(t, []int{2, 3}, platesBetweenCandles("**|**|***|", queries1))

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
	assert.Equal(t, []int{9, 0, 0, 0, 0}, platesBetweenCandles("***|**|*****|**||**|*", queries2))
}
