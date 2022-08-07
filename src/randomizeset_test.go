package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	randomizedSet := Constructor()
	assert.True(t, randomizedSet.Insert(1))
	assert.False(t, randomizedSet.Remove(2))
	assert.True(t, randomizedSet.Insert(2))
	assert.Equal(t, 2, randomizedSet.GetRandom())
	assert.True(t, randomizedSet.Remove(1))
	assert.False(t, randomizedSet.Insert(2))
	assert.Equal(t, 2, randomizedSet.GetRandom())
}
