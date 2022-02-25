package main

import (
	"reflect"
	"runtime"
	"testing"
)

func nameof(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func testTwoStringsAndReturnStringFramework(t *testing.T, testFunction func(str1, str2 string) string, str1, str2, expected string) {
	actual := testFunction(str1, str2)
	if actual != expected {
		t.Errorf("%s(%s, %s) = %s; expected %s", nameof(testFunction), str1, str2, actual, expected)
	}
}

func testArrayIntegerAndReturnArrayFramework(t *testing.T, testFunction func(array []int, target int) []int, input struct {
	array  []int
	target int
}, expected []int) {
	actual := testFunction(input.array, input.target)
	for index, val := range actual {
		if val != expected[index] {
			t.Errorf("%s(%d, %d) = %v; expected %v", nameof(testFunction), input.array, input.target, actual, expected)
		}
	}
}

func TestTwoSum(t *testing.T) {
	testArrayIntegerAndReturnArrayFramework(t, twoSum, struct {
		array  []int
		target int
	}{[]int{2, 7, 11, 15}, 9}, []int{0, 1})
	testArrayIntegerAndReturnArrayFramework(t, twoSum, struct {
		array  []int
		target int
	}{[]int{3, 2, 4}, 6}, []int{1, 2})
	testArrayIntegerAndReturnArrayFramework(t, twoSum, struct {
		array  []int
		target int
	}{[]int{3, 3}, 6}, []int{0, 1})
}

func TestComplexNumberMultiply(t *testing.T) {
	testTwoStringsAndReturnStringFramework(t, complexNumberMultiply, "1+1i", "1+1i", "0+2i")
	testTwoStringsAndReturnStringFramework(t, complexNumberMultiply, "1+-1i", "1+-1i", "0+-2i")
}
