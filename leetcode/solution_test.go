package leetcode

import (
	"datastructures"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}

func TestAddTwoNumbers(t *testing.T) {
	assert.Equal(t, addTwoNumbers(datastructures.CreateListNode([]int{2, 4, 3}), datastructures.CreateListNode([]int{5, 6, 4})), datastructures.CreateListNode([]int{7, 0, 8}))
	assert.Equal(t, addTwoNumbers(datastructures.CreateListNode([]int{0}), datastructures.CreateListNode([]int{0})), datastructures.CreateListNode([]int{0}))
	assert.Equal(t, addTwoNumbers(datastructures.CreateListNode([]int{9, 9, 9, 9, 9, 9, 9}), datastructures.CreateListNode([]int{9, 9, 9, 9})), datastructures.CreateListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, 2.00000, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.50000, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
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
	assert.Equal(t, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, pathSum(datastructures.CreateTreeNode("5,4,8,11,null,13,4,7,2,null,null,5,1"), 22))
	assert.Equal(t, [][]int(nil), pathSum(datastructures.CreateTreeNode("1,2,3"), 5))
}

func TestPreorderTraversal(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(datastructures.CreateTreeNode("1,null,2,3")))
	assert.Equal(t, []int(nil), preorderTraversal(datastructures.CreateTreeNode("")))
	assert.Equal(t, []int{1}, preorderTraversal(datastructures.CreateTreeNode("1")))
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
	assert.Equal(t, datastructures.CreateListNode([]int{5, 4, 3, 2, 1}), reverseList(datastructures.CreateListNode([]int{1, 2, 3, 4, 5})))
	assert.Equal(t, datastructures.CreateListNode([]int{2, 1}), reverseList(datastructures.CreateListNode([]int{1, 2})))
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

func TestMinAddToMakeValid(t *testing.T) {
	assert.Equal(t, 1, minAddToMakeValid("())"))
	assert.Equal(t, 3, minAddToMakeValid("((("))
}

func TestThreeEqualParts(t *testing.T) {
	assert.Equal(t, []int{0, 3}, threeEqualParts([]int{1, 0, 1, 0, 1}))
	assert.Equal(t, []int{-1, -1}, threeEqualParts([]int{1, 1, 0, 1, 1}))
	assert.Equal(t, []int{0, 2}, threeEqualParts([]int{1, 1, 0, 0, 1}))
}

func TestDistinctSubseqII(t *testing.T) {
	assert.Equal(t, 7, distinctSubseqII("abc"))
	assert.Equal(t, 6, distinctSubseqII("aba"))
	assert.Equal(t, 3, distinctSubseqII("aaa"))
}

func TestValidateStackSequences(t *testing.T) {
	assert.Equal(t, true, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	assert.Equal(t, false, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func TestOptimalDivision(t *testing.T) {
	assert.Equal(t, "1000/(100/10/2)", optimalDivision([]int{1000, 100, 10, 2}))
}

func TestPreorder(t *testing.T) {
	var root1 = &datastructures.Node{Val: 1, Children: []*datastructures.Node{}}
	var root1ChildNode3 = &datastructures.Node{Val: 3, Children: []*datastructures.Node{}}
	root1ChildNode3.Children = append(root1ChildNode3.Children, &datastructures.Node{Val: 5, Children: []*datastructures.Node{}})
	root1ChildNode3.Children = append(root1ChildNode3.Children, &datastructures.Node{Val: 6, Children: []*datastructures.Node{}})
	root1.Children = append(root1.Children, root1ChildNode3)
	root1.Children = append(root1.Children, &datastructures.Node{Val: 2, Children: []*datastructures.Node{}})
	root1.Children = append(root1.Children, &datastructures.Node{Val: 4, Children: []*datastructures.Node{}})

	var root2 = &datastructures.Node{Val: 1, Children: []*datastructures.Node{}}

	root2.Children = append(root2.Children, &datastructures.Node{Val: 2, Children: []*datastructures.Node{}})

	var root2ChildNode3 = &datastructures.Node{Val: 3, Children: []*datastructures.Node{}}
	var root2ChildNode7 = &datastructures.Node{Val: 7, Children: []*datastructures.Node{}}
	var root2ChildNode11 = &datastructures.Node{Val: 11, Children: []*datastructures.Node{}}
	root2ChildNode11.Children = append(root2ChildNode11.Children, &datastructures.Node{Val: 14, Children: []*datastructures.Node{}})
	root2ChildNode7.Children = append(root2ChildNode7.Children, root2ChildNode11)
	root2ChildNode3.Children = append(root2ChildNode3.Children, &datastructures.Node{Val: 6, Children: []*datastructures.Node{}})
	root2ChildNode3.Children = append(root2ChildNode3.Children, root2ChildNode7)

	var root2ChildNode8 = &datastructures.Node{Val: 8, Children: []*datastructures.Node{}}
	root2ChildNode8.Children = append(root2ChildNode8.Children, &datastructures.Node{Val: 12, Children: []*datastructures.Node{}})

	var root2ChildNode9 = &datastructures.Node{Val: 9, Children: []*datastructures.Node{}}
	root2ChildNode9.Children = append(root2ChildNode9.Children, &datastructures.Node{Val: 13, Children: []*datastructures.Node{}})

	var root2ChildNode4 = &datastructures.Node{Val: 4, Children: []*datastructures.Node{}}
	root2ChildNode4.Children = append(root2ChildNode4.Children, root2ChildNode8)

	var root2ChildNode5 = &datastructures.Node{Val: 5, Children: []*datastructures.Node{}}
	root2ChildNode5.Children = append(root2ChildNode5.Children, root2ChildNode9)
	root2ChildNode5.Children = append(root2ChildNode5.Children, &datastructures.Node{Val: 10, Children: []*datastructures.Node{}})

	root2.Children = append(root2.Children, root2ChildNode3)
	root2.Children = append(root2.Children, root2ChildNode4)
	root2.Children = append(root2.Children, root2ChildNode5)

	assert.Equal(t, []int{1, 3, 5, 6, 2, 4}, preorder(root1))
	assert.Equal(t, []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}, preorder(root2))
}

func TestPostorder(t *testing.T) {
	var root1 = &datastructures.Node{Val: 1, Children: []*datastructures.Node{}}
	var root1ChildNode3 = &datastructures.Node{Val: 3, Children: []*datastructures.Node{}}
	root1ChildNode3.Children = append(root1ChildNode3.Children, &datastructures.Node{Val: 5, Children: []*datastructures.Node{}})
	root1ChildNode3.Children = append(root1ChildNode3.Children, &datastructures.Node{Val: 6, Children: []*datastructures.Node{}})
	root1.Children = append(root1.Children, root1ChildNode3)
	root1.Children = append(root1.Children, &datastructures.Node{Val: 2, Children: []*datastructures.Node{}})
	root1.Children = append(root1.Children, &datastructures.Node{Val: 4, Children: []*datastructures.Node{}})

	var root2 = &datastructures.Node{Val: 1, Children: []*datastructures.Node{}}

	root2.Children = append(root2.Children, &datastructures.Node{Val: 2, Children: []*datastructures.Node{}})

	var root2ChildNode3 = &datastructures.Node{Val: 3, Children: []*datastructures.Node{}}
	var root2ChildNode7 = &datastructures.Node{Val: 7, Children: []*datastructures.Node{}}
	var root2ChildNode11 = &datastructures.Node{Val: 11, Children: []*datastructures.Node{}}
	root2ChildNode11.Children = append(root2ChildNode11.Children, &datastructures.Node{Val: 14, Children: []*datastructures.Node{}})
	root2ChildNode7.Children = append(root2ChildNode7.Children, root2ChildNode11)
	root2ChildNode3.Children = append(root2ChildNode3.Children, &datastructures.Node{Val: 6, Children: []*datastructures.Node{}})
	root2ChildNode3.Children = append(root2ChildNode3.Children, root2ChildNode7)

	var root2ChildNode8 = &datastructures.Node{Val: 8, Children: []*datastructures.Node{}}
	root2ChildNode8.Children = append(root2ChildNode8.Children, &datastructures.Node{Val: 12, Children: []*datastructures.Node{}})

	var root2ChildNode9 = &datastructures.Node{Val: 9, Children: []*datastructures.Node{}}
	root2ChildNode9.Children = append(root2ChildNode9.Children, &datastructures.Node{Val: 13, Children: []*datastructures.Node{}})

	var root2ChildNode4 = &datastructures.Node{Val: 4, Children: []*datastructures.Node{}}
	root2ChildNode4.Children = append(root2ChildNode4.Children, root2ChildNode8)

	var root2ChildNode5 = &datastructures.Node{Val: 5, Children: []*datastructures.Node{}}
	root2ChildNode5.Children = append(root2ChildNode5.Children, root2ChildNode9)
	root2ChildNode5.Children = append(root2ChildNode5.Children, &datastructures.Node{Val: 10, Children: []*datastructures.Node{}})

	root2.Children = append(root2.Children, root2ChildNode3)
	root2.Children = append(root2.Children, root2ChildNode4)
	root2.Children = append(root2.Children, root2ChildNode5)

	assert.Equal(t, []int{5, 6, 3, 2, 4, 1}, postorder(root1))
	assert.Equal(t, []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1}, postorder(root2))
}

func TestFindRestaurant(t *testing.T) {
	assert.Equal(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	assert.Equal(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
}

func TestExclusiveTime(t *testing.T) {
	assert.Equal(t, []int{3, 4}, exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
	assert.Equal(t, []int{8}, exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	assert.Equal(t, []int{7, 1}, exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}))
	assert.Equal(t, []int{8, 1}, exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}))
	assert.Equal(t, []int{1}, exclusiveTime(1, []string{"0:start:0", "0:end:0"}))
}

func TestFindLongestChain(t *testing.T) {
	assert.Equal(t, 2, findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	assert.Equal(t, 3, findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))
}

func TestFindDuplicateSubtrees(t *testing.T) {
	assert.Equal(t, []*datastructures.TreeNode{
		datastructures.CreateTreeNode("4"),
		datastructures.CreateTreeNode("2,4"),
	}, findDuplicateSubtrees(datastructures.CreateTreeNode("1,2,3,4,null,2,4,null,null,4")))
	assert.Equal(t, []*datastructures.TreeNode{
		datastructures.CreateTreeNode("1"),
	}, findDuplicateSubtrees(datastructures.CreateTreeNode("2,1,1")))
	assert.Equal(t, []*datastructures.TreeNode{
		datastructures.CreateTreeNode("3"),
		datastructures.CreateTreeNode("2,3"),
	}, findDuplicateSubtrees(datastructures.CreateTreeNode("2,2,2,3,null,3,null")))
}

func TestPrintTree(t *testing.T) {
	assert.Equal(t, [][]string{{"", "1", ""}, {"2", "", ""}}, printTree(datastructures.CreateTreeNode("1,2")))
	assert.Equal(t, [][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "3", ""}, {"", "", "4", "", "", "", ""}}, printTree(datastructures.CreateTreeNode("1,2,3,null,4")))
}

func TestFindClosestElements(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	assert.Equal(t, []int{1, 2, 3, 4}, findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
}

func TestWidthOfBinaryTree(t *testing.T) {
	assert.Equal(t, 4, widthOfBinaryTree(datastructures.CreateTreeNode("1,3,2,5,3,null,9")))
	assert.Equal(t, 7, widthOfBinaryTree(datastructures.CreateTreeNode("1,3,2,5,null,null,9,6,null,7")))
	assert.Equal(t, 2, widthOfBinaryTree(datastructures.CreateTreeNode("1,3,2,5")))
}

func TestConstructArray(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, constructArray(3, 1))
	assert.Equal(t, []int{1, 3, 2}, constructArray(3, 2))
}

func TestTrimBST(t *testing.T) {
	assert.Equal(t, datastructures.CreateTreeNode("1,null,2"), trimBST(datastructures.CreateTreeNode("1,0,2"), 1, 2))
	assert.Equal(t, datastructures.CreateTreeNode("3,2,null,1"), trimBST(datastructures.CreateTreeNode("3,0,4,null,2,null,null,1"), 1, 3))
}

func TestMaximumSwap(t *testing.T) {
	assert.Equal(t, 7236, maximumSwap(2736))
	assert.Equal(t, 9973, maximumSwap(9973))
}

func TestFlipLights(t *testing.T) {
	assert.Equal(t, 2, flipLights(1, 1))
	assert.Equal(t, 3, flipLights(2, 1))
	assert.Equal(t, 4, flipLights(3, 1))
}

func TestLongestUnivaluePath(t *testing.T) {
	assert.Equal(t, 2, longestUnivaluePath(datastructures.CreateTreeNode("5,4,5,1,1,null,5")))
	assert.Equal(t, 2, longestUnivaluePath(datastructures.CreateTreeNode("1,4,5,4,4,null,5")))
}

func TestPivotIndex(t *testing.T) {
	assert.Equal(t, 3, pivotIndex([]int{2, 3, -1, 8, 4}))
	assert.Equal(t, 2, pivotIndex([]int{1, -1, 4}))
	assert.Equal(t, -1, pivotIndex([]int{2, 5}))
}

func TestSelfDividingNumbers(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}, selfDividingNumbers(1, 22))
	assert.Equal(t, []int{48, 55, 66, 77}, selfDividingNumbers(48, 85))
}

func TestMaxChunksToSorted(t *testing.T) {
	assert.Equal(t, 1, maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	assert.Equal(t, 4, maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}

func TestCanTransform(t *testing.T) {
	assert.Equal(t, true, canTransform("RXXLRXRXL", "XRLXXRRLX"))
	assert.Equal(t, false, canTransform("X", "L"))
}

func TestKthGrammar(t *testing.T) {
	assert.Equal(t, 0, kthGrammar(1, 1))
	assert.Equal(t, 0, kthGrammar(2, 1))
	assert.Equal(t, 1, kthGrammar(2, 2))
}

func TestPreimageSizeFZF(t *testing.T) {
	assert.Equal(t, 5, preimageSizeFZF(0))
	assert.Equal(t, 0, preimageSizeFZF(5))
	assert.Equal(t, 5, preimageSizeFZF(3))
}

func TestMinSwap(t *testing.T) {
	assert.Equal(t, 1, minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}))
	assert.Equal(t, 1, minSwap([]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}))
}

func TestUniqueMorseRepresentations(t *testing.T) {
	assert.Equal(t, 2, uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	assert.Equal(t, 1, uniqueMorseRepresentations([]string{"a"}))
}

func TestNumberOfLines(t *testing.T) {
	assert.Equal(t, []int{3, 60}, numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	assert.Equal(t, []int{2, 4}, numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func TestSubdomainVisits(t *testing.T) {
	expected1 := []string{"9001 leetcode.com", "9001 com", "9001 discuss.leetcode.com"}
	expected2 := []string{"901 mail.com", "951 com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org", "5 org", "900 google.mail.com"}
	actual1 := subdomainVisits([]string{"9001 discuss.leetcode.com"})
	actual2 := subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"})
	sort.Strings(expected1)
	sort.Strings(actual1)
	sort.Strings(expected2)
	sort.Strings(actual2)
	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

func TestNumComponents(t *testing.T) {
	assert.Equal(t, 2, numComponents(datastructures.CreateListNode([]int{0, 1, 2, 3}), []int{0, 1, 3}))
	assert.Equal(t, 2, numComponents(datastructures.CreateListNode([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}))
}

func TestMostCommonWord(t *testing.T) {
	assert.Equal(t, "ball", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

func TestUniqueLetterString(t *testing.T) {
	assert.Equal(t, 10, uniqueLetterString("ABC"))
	assert.Equal(t, 8, uniqueLetterString("ABA"))
	assert.Equal(t, 92, uniqueLetterString("LEETCODE"))
}

func TestScoreOfParentheses(t *testing.T) {
	assert.Equal(t, 1, scoreOfParentheses("()"))
	assert.Equal(t, 2, scoreOfParentheses("(())"))
	assert.Equal(t, 2, scoreOfParentheses("()()"))
}

func TestMinCostToHireWorkers(t *testing.T) {
	assert.Equal(t, 105.00000, round(minCostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2), 5))
	assert.Equal(t, 30.66667, round(minCostToHireWorkers([]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3), 5))
}

func TestAdvantageCount(t *testing.T) {
	assert.Equal(t, []int{2, 11, 7, 15}, advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	assert.Equal(t, []int{24, 32, 8, 12}, advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}

func TestProjectionArea(t *testing.T) {
	assert.Equal(t, 17, projectionArea([][]int{{1, 2}, {3, 4}}))
	assert.Equal(t, 5, projectionArea([][]int{{2}}))
	assert.Equal(t, 8, projectionArea([][]int{{1, 0}, {0, 2}}))
}

func TestPossibleBipartition(t *testing.T) {
	assert.Equal(t, true, possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	assert.Equal(t, false, possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	assert.Equal(t, false, possibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))
}

func TestTotalFruit(t *testing.T) {
	assert.Equal(t, 3, totalFruit([]int{1, 2, 1}))
	assert.Equal(t, 3, totalFruit([]int{0, 1, 2, 2}))
	assert.Equal(t, 4, totalFruit([]int{1, 2, 3, 2, 2}))
}

func TestPartitionDisjoint(t *testing.T) {
	assert.Equal(t, 3, partitionDisjoint([]int{5, 0, 3, 8, 6}))
	assert.Equal(t, 4, partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
}

func TestRepeatedNTimes(t *testing.T) {
	assert.Equal(t, 3, repeatedNTimes([]int{1, 2, 3, 3}))
	assert.Equal(t, 2, repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	assert.Equal(t, 5, repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}

func TestInsertIntoMaxTree(t *testing.T) {
	assert.Equal(t, datastructures.CreateTreeNode("5,4,null,1,3,null,null,2"),
		insertIntoMaxTree(datastructures.CreateTreeNode("4,1,3,null,null,2"), 5))
	assert.Equal(t, datastructures.CreateTreeNode("5,2,4,null,1,null,3"),
		insertIntoMaxTree(datastructures.CreateTreeNode("5,2,4,null,1"), 3))
	assert.Equal(t, datastructures.CreateTreeNode("5,2,4,null,1,3"),
		insertIntoMaxTree(datastructures.CreateTreeNode("5,2,3,null,1"), 4))
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

func TestBuildArray(t *testing.T) {
	assert.Equal(t, []string{"Push", "Push", "Pop", "Push"}, buildArray([]int{1, 3}, 3))
	assert.Equal(t, []string{"Push", "Push", "Push"}, buildArray([]int{1, 2, 3}, 3))
	assert.Equal(t, []string{"Push", "Push"}, buildArray([]int{1, 2}, 4))
}

func TestBusyStudent(t *testing.T) {
	assert.Equal(t, 1, busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
	assert.Equal(t, 1, busyStudent([]int{4}, []int{4}, 4))
}

func TestIsPrefixOfWord(t *testing.T) {
	assert.Equal(t, 4, isPrefixOfWord("i love eating burger", "burg"))
	assert.Equal(t, 2, isPrefixOfWord("this problem is an easy problem", "pro"))
	assert.Equal(t, -1, isPrefixOfWord("i am tired", "you"))
}

func TestMaxProduct(t *testing.T) {
	assert.Equal(t, 12, maxProduct([]int{3, 4, 5, 2}))
	assert.Equal(t, 16, maxProduct([]int{1, 5, 4, 5}))
	assert.Equal(t, 12, maxProduct([]int{3, 7}))
}

func TestShuffle(t *testing.T) {
	assert.Equal(t, []int{2, 3, 5, 4, 1, 7}, shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	assert.Equal(t, []int{1, 4, 2, 3, 3, 2, 4, 1}, shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	assert.Equal(t, []int{1, 2, 1, 2}, shuffle([]int{1, 1, 2, 2}, 2))
}

func TestFinalPrices(t *testing.T) {
	assert.Equal(t, []int{4, 2, 4, 2, 3}, finalPrices([]int{8, 4, 6, 2, 3}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, finalPrices([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{9, 0, 1, 6}, finalPrices([]int{10, 1, 1, 6}))
}

func TestNumSpecial(t *testing.T) {
	assert.Equal(t, 1, numSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))
	assert.Equal(t, 3, numSpecial([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
}

func TestReorderSpaces(t *testing.T) {
	assert.Equal(t, "this   is   a   sentence", reorderSpaces("  this   is  a sentence "))
	assert.Equal(t, "practice   makes   perfect ", reorderSpaces(" practice   makes   perfect"))
}

func TestMinOperations(t *testing.T) {
	assert.Equal(t, 2, minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
	assert.Equal(t, 3, minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
	assert.Equal(t, 0, minOperations([]string{"d1/", "../", "../", "../"}))
}

func TestSpecialArray(t *testing.T) {
	assert.Equal(t, 2, specialArray([]int{3, 5}))
	assert.Equal(t, -1, specialArray([]int{0, 0}))
	assert.Equal(t, 3, specialArray([]int{0, 4, 3, 0, 4}))
}

func TestTrimMean(t *testing.T) {
	assert.True(t, (trimMean([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3})-2.00000) <= 0.00001)
	assert.True(t, (trimMean([]int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0})-4.00000) <= 0.00001)
	assert.True(t, (trimMean([]int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4})-4.77778) <= 0.00001)
}

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	assert.Equal(t, 0, maxLengthBetweenEqualCharacters("aa"))
	assert.Equal(t, 2, maxLengthBetweenEqualCharacters("abca"))
	assert.Equal(t, -1, maxLengthBetweenEqualCharacters("bczyx"))
}

func TestFrequencySort(t *testing.T) {
	assert.Equal(t, []int{3, 1, 1, 2, 2, 2}, frequencySort([]int{1, 1, 2, 2, 2, 3}))
	assert.Equal(t, []int{1, 3, 3, 2, 2}, frequencySort([]int{2, 3, 1, 3, 2}))
	assert.Equal(t, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}, frequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}

func TestReformatNumber(t *testing.T) {
	assert.Equal(t, "123-456", reformatNumber("1-23-45 6"))
	assert.Equal(t, "123-45-67", reformatNumber("123 4-567"))
	assert.Equal(t, "123-456-78", reformatNumber("123 4-5678"))
}

func TestMergeAlternately(t *testing.T) {
	assert.Equal(t, "apbqcr", mergeAlternately("abc", "pqr"))
	assert.Equal(t, "apbqrs", mergeAlternately("ab", "pqrs"))
	assert.Equal(t, "apbqcd", mergeAlternately("abcd", "pq"))
}

func TestCheckOnesSegment(t *testing.T) {
	assert.Equal(t, false, checkOnesSegment("1001"))
	assert.Equal(t, true, checkOnesSegment("110"))
}

func TestAreAlmostEqual(t *testing.T) {
	assert.Equal(t, true, areAlmostEqual("bank", "kanb"))
	assert.Equal(t, false, areAlmostEqual("attack", "defend"))
	assert.Equal(t, true, areAlmostEqual("kelb", "kelb"))
}

func TestMaxAscendingSum(t *testing.T) {
	assert.Equal(t, 65, maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
	assert.Equal(t, 150, maxAscendingSum([]int{10, 20, 30, 40, 50}))
	assert.Equal(t, 33, maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}

func TestFindTheWinner(t *testing.T) {
	assert.Equal(t, 3, findTheWinner(5, 2))
	assert.Equal(t, 1, findTheWinner(6, 5))
}

func TestFindMiddleIndex(t *testing.T) {
	assert.Equal(t, 3, findMiddleIndex([]int{2, 3, -1, 8, 4}))
	assert.Equal(t, 2, findMiddleIndex([]int{1, -1, 4}))
	assert.Equal(t, -1, findMiddleIndex([]int{2, 5}))
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
	assert.Equal(t, []int{2, 3}, platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
	assert.Equal(t, []int{9, 0, 0, 0, 0}, platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}))
}
