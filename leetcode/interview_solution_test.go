package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	assert.Equal(t, true, checkPermutation("abc", "bca"))
	assert.Equal(t, false, checkPermutation("abc", "bad"))
}
