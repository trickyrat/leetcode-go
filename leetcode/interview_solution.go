package leetcode

// 01.02. Check Permutation LCCI
func checkPermutation(s1, s2 string) bool {
	var c1, c2 [128]int
	for _, ch := range s1 {
		c1[ch]++
	}
	for _, ch := range s2 {
		c2[ch]++
	}
	return c1 == c2
}
