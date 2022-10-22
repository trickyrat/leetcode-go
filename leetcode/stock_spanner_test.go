package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStockSpanner_Next(t *testing.T) {
	stockSpanner := StockSpannerConstructor()
	assert.Equal(t, 1, stockSpanner.Next(100))
	assert.Equal(t, 1, stockSpanner.Next(80))
	assert.Equal(t, 1, stockSpanner.Next(60))
	assert.Equal(t, 2, stockSpanner.Next(70))
	assert.Equal(t, 1, stockSpanner.Next(60))
	assert.Equal(t, 4, stockSpanner.Next(75))
	assert.Equal(t, 6, stockSpanner.Next(85))
}
