package main

import "testing"

func testBoolEqualBase(t *testing.T, expected bool, actual bool) {
	if actual != expected {
		t.Errorf("actual: %t, but expected: %t", actual, expected)
	}
}

func testIntEqualBase(t *testing.T, expected []int, actual int) {
	isExist := false
	for _, item := range expected {
		if item == actual {
			isExist = true
			break
		}
	}
	if !isExist {
		t.Errorf("actual: %d, but expected: %d", actual, expected)
	}
}

func testRandomizedSetInsert(t *testing.T, set *RandomizedSet, input int, expected bool) {
	actual := set.Insert(input)
	testBoolEqualBase(t, expected, actual)
}

func testRandomizedSetRemove(t *testing.T, set *RandomizedSet, input int, expected bool) {
	actual := set.Remove(input)
	testBoolEqualBase(t, expected, actual)
}

func testRandomizedSetGetRandom(t *testing.T, set *RandomizedSet, expected []int) {
	actual := set.GetRandom()
	testIntEqualBase(t, expected, actual)
}

func TestRandomizedSet(t *testing.T) {
	randomizedSet := Constructor()
	testRandomizedSetInsert(t, &randomizedSet, 1, true)
	testRandomizedSetRemove(t, &randomizedSet, 2, false)
	testRandomizedSetInsert(t, &randomizedSet, 2, true)
	testRandomizedSetGetRandom(t, &randomizedSet, []int{1, 2})
	testRandomizedSetRemove(t, &randomizedSet, 1, true)
	testRandomizedSetInsert(t, &randomizedSet, 2, false)
	testRandomizedSetGetRandom(t, &randomizedSet, []int{2})
}
