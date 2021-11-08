package leetcode


// 191. Number of 1 Bits
func hammingWeight(num uint32) int {
	var res int = 0
	for  {
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
	return n > 0 && (n & (n - 1)) == 0
}
