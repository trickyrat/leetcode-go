package leetcode

import (
	"container/heap"
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

// 646. Maximum Length of Pair Chain
func findLongestChain(pairs [][]int) (res int) {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	curr := math.MinInt32
	for _, p := range pairs {
		if curr < p[0] {
			curr = p[1]
			res++
		}
	}
	return
}

// 652. Find Duplicate Subtrees
func findDuplicateSubtrees(root *datastructures.TreeNode) []*datastructures.TreeNode {
	type pair struct {
		node  *datastructures.TreeNode
		index int
	}
	repeat := map[*datastructures.TreeNode]struct{}{}
	seen := map[[3]int]pair{}
	index := 0
	var dfs func(*datastructures.TreeNode) int
	dfs = func(node *datastructures.TreeNode) int {
		if node == nil {
			return 0
		}
		triple := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
		if p, ok := seen[triple]; ok {
			repeat[p.node] = struct{}{}
			return p.index
		}
		index++
		seen[triple] = pair{node, index}
		return index
	}
	dfs(root)
	res := make([]*datastructures.TreeNode, 0, len(repeat))
	for node := range repeat {
		res = append(res, node)
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

// 658. Find K Closest Elements
func findClosestElements(arr []int, k, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	n := len(arr)
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= n || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}

// 662. Maximum Width of Binary Tree
func widthOfBinaryTree(root *datastructures.TreeNode) int {
	levelMin := map[int]int{}
	var dfs func(*datastructures.TreeNode, int, int) int
	dfs = func(node *datastructures.TreeNode, depth, index int) int {
		if node == nil {
			return 0
		}
		if _, ok := levelMin[depth]; !ok {
			levelMin[depth] = index
		}
		return max(index-levelMin[depth]+1,
			max(
				dfs(node.Left, depth+1, index*2),
				dfs(node.Right, depth+1, index*2+1),
			))
	}
	return dfs(root, 1, 1)
}

// 667. Beautiful Arrangement II
func constructArray(n, k int) []int {
	res := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		res = append(res, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		res = append(res, i)
		if i != j {
			res = append(res, j)
		}
		j--
	}
	return res
}

// 669. Trim a Binary Search Tree
func trimBST(root *datastructures.TreeNode, low, high int) *datastructures.TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	for node := root; node.Left != nil; {
		if node.Left.Val < low {
			node.Left = node.Left.Right
		} else {
			node = node.Left
		}
	}
	for node := root; node.Right != nil; {
		if node.Right.Val > high {
			node.Right = node.Right.Left
		} else {
			node = node.Right
		}
	}
	return root
}

// 670. Maximum Swap
func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)
	maxIndex, index1, index2 := n-1, -1, -1
	for i := n - 1; i >= 0; i-- {
		if s[i] > s[maxIndex] {
			maxIndex = i
		} else if s[i] < s[maxIndex] {
			index1, index2 = i, maxIndex
		}
	}
	if index1 < 0 {
		return num
	}
	s[index1], s[index2] = s[index2], s[index1]
	v, _ := strconv.Atoi(string(s))
	return v
}

// 687. Longest Univalue Path
func longestUnivaluePath(root *datastructures.TreeNode) (res int) {
	var dfs func(node *datastructures.TreeNode) int
	dfs = func(node *datastructures.TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		left1, right1 := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			left1 = left + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right1 = right + 1
		}
		res = max(res, left1+right1)
		return max(left1, right1)
	}
	dfs(root)
	return res
}

// 724. Find Pivot Index
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

// 728.Self Dividing Numbers
func selfDividingNumbers(left, right int) (ans []int) {
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ans = append(ans, i)
		}
	}
	return
}

// 793. Preimage Size of Factorial Zeroes Function
func preimageSizeFZF(k int) int {
	var zeta func(n int) int
	zeta = func(n int) (res int) {
		for n > 0 {
			n /= 5
			res += n
		}
		return res
	}
	var nx func(n int) int
	nx = func(n int) int {
		return sort.Search(5*n, func(x int) bool { return zeta(x) >= n })
	}
	return nx(k+1) - nx(k)
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

// 828. Count Unique Characters of All Substrings of a Given String
func uniqueLetterString(s string) (res int) {
	index := map[rune][]int{}
	for i, c := range s {
		index[c] = append(index[c], i)
	}
	for _, arr := range index {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			res += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 857. Minimum Cost to Hire K Workers
func minCostToHireWorkers(quality, wage []int, k int) float64 {
	n := len(quality)
	hire := make([]int, n)
	for i := range hire {
		hire[i] = i
	}
	sort.Slice(hire, func(i, j int) bool {
		a, b := hire[i], hire[j]
		return quality[a]*wage[b] > quality[b]*wage[a]
	})
	totalQuality := 0
	queue := hp{}
	for i := 0; i < k-1; i++ {
		totalQuality += quality[hire[i]]
		heap.Push(&queue, quality[hire[i]])
	}
	res := 1e9
	for i := k - 1; i < n; i++ {
		index := hire[i]
		totalQuality += quality[index]
		heap.Push(&queue, quality[index])
		res = math.Min(res, float64(wage[index])/float64(quality[index])*float64(totalQuality))
		totalQuality -= heap.Pop(&queue).(int)
	}
	return res
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

// 946. Validate Stack Sequences
func validateStackSequences(pushed, popped []int) bool {
	var stack []int
	j := 0
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
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

// 998. Maximum Binary Tree II
func insertIntoMaxTree(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	var parent *datastructures.TreeNode
	for curr := root; curr != nil; curr = curr.Right {
		if val > curr.Val {
			if parent == nil {
				return &datastructures.TreeNode{Val: val, Left: root}
			}
			parent.Right = &datastructures.TreeNode{Val: val, Left: curr}
			return root
		}
		parent = curr
	}
	parent.Right = &datastructures.TreeNode{Val: val}
	return root
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

// 1464. Maximum Product of Two Elements in an Array
func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if a < b {
		a, b = b, a
	}
	for _, num := range nums[2:] {
		if num > a {
			a, b = num, a
		} else if num > b {
			b = num
		}
	}
	return (a - 1) * (b - 1)
}

// 1470. Shuffle the Array
func shuffle(nums []int, n int) []int {
	res := make([]int, n*2)
	for i, num := range nums[:n] {
		res[2*i] = num
		res[2*i+1] = nums[i+n]
	}
	return res
}

// 1475. Final Prices With a Special Discount in a Shop
func finalPrices(prices []int) []int {
	n := len(prices)
	res := make([]int, n)
	stack := []int{0}
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for len(stack) > 1 && stack[len(stack)-1] > p {
			stack = stack[:len(stack)-1]
		}
		res[i] = p - stack[len(stack)-1]
		stack = append(stack, p)
	}
	return res
}

// 1582. Special Positions in a Binary Matrix
func numSpecial(mat [][]int) (res int) {
	for i, row := range mat {
		count := 0
		for _, x := range row {
			count += x
		}
		if i == 0 {
			count--
		}
		if count > 0 {
			for j, x := range row {
				if x == 1 {
					mat[0][j] += count
				}
			}
		}
	}
	for _, x := range mat[0] {
		if x == 1 {
			res++
		}
	}
	return
}

// 1592. Rearrange Spaces Between Words
func reorderSpaces(text string) string {
	words := strings.Fields(text)
	space := strings.Count(text, " ")
	lengthOfWord := len(words) - 1
	if lengthOfWord == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/lengthOfWord)) + strings.Repeat(" ", space%lengthOfWord)
}

// 1598. Crawler Log Folder
func minOperations(logs []string) int {
	depth := 0
	for _, log := range logs {
		if log == "./" {
			continue
		} else if log == "../" {
			if depth > 0 {
				depth--
			}
		} else {
			depth++
		}
	}
	return depth
}

// 1608. Special Array With X Elements Greater Than or Equal X
func specialArray(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}

// 1619. Mean of Array After Removing Some Elements
func trimMean(arr []int) float64 {
	n := len(arr)
	sort.Ints(arr)
	sum := 0
	for _, i := range arr[n/20 : (19 * n / 20)] {
		sum += i
	}
	return float64(sum*10) / float64(n*9)
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
func findMiddleIndex(nums []int) int {
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
