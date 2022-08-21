package datastructures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNodeToString(t *testing.T) {
	head1 := CreateListNode([]int{1, 2, 3, 4})
	head2 := CreateListNode([]int{-6, 0, 6})
	assert.Equal(t, "1->2->3->4", head1.ToString())
	assert.Equal(t, "-6->0->6", head2.ToString())
}
