package leetcode

import "math"

type StockSpanner struct {
	stack [][2]int
	index int
}

func StockSpannerConstructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, -1}
}

func (s *StockSpanner) Next(price int) int {
	s.index++
	for price >= s.stack[len(s.stack)-1][1] {
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{s.index, price})
	return s.index - s.stack[len(s.stack)-2][0]
}
