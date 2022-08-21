package leetcode

import (
	"datastructures"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

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

// 2.Add Two Numbers
func addTwoNumbers(l1, l2 *datastructures.ListNode) (head *datastructures.ListNode) {
	var tail *datastructures.ListNode
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
			head = &datastructures.ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &datastructures.ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &datastructures.ListNode{Val: carry}
	}
	return
}

// 6.Zigzag Conversion
func zconvert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := 2*r - 2
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ {
		for j := 0; j < n-i; j += t {
			ans = append(ans, s[j+i])
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i])
			}
		}
	}
	return string(ans)
}

// 7.Reverse Integer
func reverseInt(x int) (res int) {
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
	return
}

// 113.Path Sum II
func pathSum(root *datastructures.TreeNode, targetSum int) (ans [][]int) {
	var path []int
	var dfs func(*datastructures.TreeNode, int)
	dfs = func(node *datastructures.TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, targetSum)
	return
}

// 144.Binary Tree Preorder Traversal
func preorderTraversal(root *datastructures.TreeNode) (res []int) {
	var stack []*datastructures.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

// 191.Number of 1 Bits
func hammingWeight(num uint32) int {
	var res = 0
	for {
		if num == 0 {
			break
		}
		num &= num - 1
		res++
	}
	return res
}

// 231.Power of Two
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// 206.Reverse Linked List
func reverseList(head *datastructures.ListNode) *datastructures.ListNode {
	var prev *datastructures.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 357.Count Numbers with Unique Digits
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	ans, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		ans += cur
	}
	return ans
}

// 386.Lexicographical Numbers
func lexicalOrder(n int) []int {
	ret := make([]int, n)
	num := 1
	for i := range ret {
		ret[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ret
}

// 393.UTF-8 Validation
func validUtf8(data []int) bool {
	for index, m := 0, len(data); index < m; {
		n := getBytes(data[index])
		if n < 0 || index+n > m {
			return false
		}
		for _, ch := range data[index+1 : index+n] {
			if ch&mask2 != mask1 {
				return false
			}
		}
		index += n
	}
	return true
}

// 467.Unique Substrings in Wraparound String
func findSubstringWraparoundString(p string) (ans int) {
	dp := [26]int{}
	k := 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 {
			k++
		} else {
			k = 1
		}
		dp[ch-'a'] = max(dp[ch-'a'], k)
	}
	for _, v := range dp {
		ans += v
	}
	return
}

// 504.Base 7
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num < 0
	if negative {
		num = -num
	}
	var s []byte
	for num > 0 {
		s = append(s, '0'+byte(num%7))
		num /= 7
	}
	if negative {
		s = append(s, '-')
	}
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return string(s)

}

// 521.Longest Uncommon Subsequence I
func findLUSLength(a, b string) (ans int) {
	if a == b {
		ans = -1
	} else {
		ans = max(len(a), len(b))
	}
	return
}

// 537.Complex Number Multiplication
func complexNumberMultiply(num1, num2 string) string {
	real1, imag1 := parseComplexNumber(num1)
	real2, imag2 := parseComplexNumber(num2)
	return fmt.Sprintf("%d+%di", real1*real2-imag1*imag2, real1*imag2+imag1*real2)
}

// 553.Optimal Division
func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	ans := &strings.Builder{}
	ans.WriteString(fmt.Sprintf("%d/(%d", nums[0], nums[1]))
	for _, num := range nums[2:] {
		ans.WriteByte('/')
		ans.WriteString(strconv.Itoa(num))
	}
	ans.WriteByte(')')
	return ans.String()
}

// 589.N-ary Tree Preorder Traversal
func preorder(root *datastructures.Node) (ans []int) {
	var dfs func(*datastructures.Node)
	dfs = func(node *datastructures.Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return
}

// 590.N-ary Tree Postorder Traversal
func postorder(root *datastructures.Node) (ans []int) {
	var dfs func(*datastructures.Node)
	dfs = func(node *datastructures.Node) {
		if node == nil {
			return
		}
		for _, ch := range node.Children {
			dfs(ch)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}

// 599.Minimum Index Sum of Two Lists
func findRestaurant(list1, list2 []string) (ans []string) {
	index := make(map[string]int, len(list1))
	for i, s := range list1 {
		index[s] = i
	}
	indexSum := math.MaxInt32
	for i, s := range list2 {
		if j, ok := index[s]; ok {
			if i+j < indexSum {
				indexSum = i + j
				ans = []string{s}
			} else if i+j == indexSum {
				ans = append(ans, s)
			}
		}
	}
	return
}

// 636. Exclusive Time of Functions
func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	type pair struct {
		index, timestamp int
	}
	var stack []pair
	for _, log := range logs {
		sp := strings.Split(log, ":")
		index, _ := strconv.Atoi(sp[0])
		timestamp, _ := strconv.Atoi(sp[2])
		if sp[1][0] == 's' { // 0:start:1
			if len(stack) > 0 {
				res[stack[len(stack)-1].index] += timestamp - stack[len(stack)-1].timestamp
			}
			stack = append(stack, pair{index, timestamp})
		} else { // 0:end:1
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[p.index] += timestamp - p.timestamp + 1
			if len(stack) > 0 {
				stack[len(stack)-1].timestamp = timestamp + 1
			}
		}
	}
	return res
}

// 655. Print Binary Tree
func printTree(root *datastructures.TreeNode) [][]string {
	height := calculateDepth(root)
	m := height + 1
	n := 1<<m - 1
	res := make([][]string, m)
	for i := range res {
		res[i] = make([]string, n)
	}
	type entry struct {
		node        *datastructures.TreeNode
		row, column int
	}
	queue := []entry{{root, 0, (n - 1) / 2}}
	for len(queue) > 0 {
		ele := queue[0]
		queue = queue[1:]
		node, row, column := ele.node, ele.row, ele.column
		res[row][column] = strconv.Itoa(node.Val)
		if node.Left != nil {
			queue = append(queue, entry{node.Left, row + 1, column - 1<<(height-row-1)})
		}
		if node.Right != nil {
			queue = append(queue, entry{node.Right, row + 1, column + 1<<(height-row-1)})
		}
	}
	return res
}

// 728.Self Dividing Numbers
func selfDividingNumbers(left, right int) (ans []int) {
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ans = append(ans, i)
		}
	}
	return
}

// 804.Unique Morse Code Words
func uniqueMorseRepresentations(words []string) int {
	var morse = []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
		"....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-",
		"...-", ".--", "-..-", "-.--", "--..",
	}
	set := map[string]struct{}{}
	for _, word := range words {
		trans := &strings.Builder{}
		for _, ch := range word {
			trans.WriteString(morse[ch-'a'])
		}
		set[trans.String()] = struct{}{}
	}
	return len(set)
}

// 806.Number of Lines To Write String
func numberOfLines(widths []int, s string) (ans []int) {
	const MaxWidth = 100
	lines, width := 1, 0
	for _, ch := range s {
		need := widths[ch-'a']
		width += need
		if width > MaxWidth {
			width = need
			lines++
		}
	}
	ans = append(ans, lines, width)
	return
}

// 819.Most Common Word
func mostCommonWord(paragraph string, banned []string) string {
	ban := map[string]bool{}
	for _, s := range banned {
		ban[s] = true
	}
	freq := map[string]int{}
	maxFreq := 0
	var word []byte
	for i, n := 0, len(paragraph); i <= n; i++ {
		if i < n && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if word != nil {
			s := string(word)
			if !ban[s] {
				freq[s]++
				maxFreq = max(maxFreq, freq[s])
			}
			word = nil
		}
	}
	for s, f := range freq {
		if f == maxFreq {
			return s
		}
	}
	return ""
}

// 883.Projection Area of 3D Shapes
func projectionArea(grid [][]int) int {
	var xyArea, yzArea, zxArea int
	for i, row := range grid {
		yzHeight, zxHeight := 0, 0
		for j, v := range row {
			if v > 0 {
				xyArea++
			}
			yzHeight = max(yzHeight, grid[j][i])
			zxHeight = max(zxHeight, v)
		}
		yzArea += yzHeight
		zxArea += zxHeight
	}
	return xyArea + yzArea + zxArea
}

// 917.Reverse Only Letters
func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	left, right := 0, len(s)-1
	for {
		for left < right && !unicode.IsLetter(rune(s[left])) {
			left++
		}
		for right > left && !unicode.IsLetter(rune(s[right])) {
			right--
		}
		if left >= right {
			break
		}
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return string(ans)
}

// 961.N-Repeated Element in Size 2N Array
func repeatedNTimes(nums []int) int {
	found := map[int]bool{}
	for _, num := range nums {
		if found[num] {
			return num
		}
		found[num] = true
	}
	return -1
}

// 1403.Minimum Subsequence in Non-Increasing Order
func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	total := 0
	for _, num := range nums {
		total += num
	}
	for i, curr := 0, 0; ; i++ {
		curr += nums[i]
		if total-curr < curr {
			return nums[:i+1]
		}
	}
}

// 1408.String Matching in an Array
func stringMatching(words []string) (res []string) {
	for i, x := range words {
		for j, y := range words {
			if j != i && strings.Contains(y, x) {
				res = append(res, x)
				break
			}
		}
	}
	return
}

// 1450. Number of Students Doing Homework at a Given Time
func busyStudent(startTime []int, endTime []int, queryTime int) (res int) {
	for i, s := range startTime {
		if s <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return
}

// 1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence
func isPrefixOfWord(sentence string, searchWord string) int {
	for i, index, n := 0, 1, len(sentence); i < n; i++ {
		start := i
		for i < n && sentence[i] != ' ' {
			i++
		}
		end := i
		if strings.HasPrefix(sentence[start:end], searchWord) {
			return index
		}
		index++
	}
	return -1
}

// 1823.Find the Winner of the Circular Game
func findTheWinner(n, k int) int {
	winner := 1
	for i := 2; i <= n; i++ {
		winner = (winner+k-1)%i + 1
	}
	return winner
}

// 1991.Find the Middle Index in Array
func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		if 2*sum+num == total {
			return i
		}
		sum += num
	}
	return -1
}

// 2006.Count Number of Pairs With Absolute Difference K
func countKDifference(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		ans += cnt[num-k] + cnt[num+k]
		cnt[num]++
	}
	return
}

// 2016.Maximum Difference Between Increasing Elements
func maximumDifference(nums []int) int {
	ans := -1
	for i, preMin := 1, nums[0]; i < len(nums); i++ {
		if nums[i] > preMin {
			ans = max(ans, nums[i]-preMin)
		} else {
			preMin = nums[i]
		}
	}
	return ans
}

// 2044.Count Number of Maximum Bitwise-OR Subsets
func countMaxOrSubsets(nums []int) (ans int) {
	maxOr := 0
	var dfs func(int, int)
	dfs = func(pos, or int) {
		if pos == len(nums) {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos])
		dfs(pos+1, or)
	}
	dfs(0, 0)
	return
}

// 2055.Plates Between Candles
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	preSum := make([]int, n)
	left := make([]int, n)
	sum, l := 0, -1
	for i, ch := range s {
		if ch == '*' {
			sum++
		} else {
			l = i
		}
		preSum[i] = sum
		left[i] = l
	}
	right := make([]int, n)
	for i, r := n-1, -1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			ans[i] = preSum[y] - preSum[x]
		}
	}
	return ans
}
