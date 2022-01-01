package leetcode

import (
	"testing"
)

type testData struct {
	name      string
	nums      []int
	target    int
	expResult []int
}

var (
	testCases []*testData = []*testData{
		{
			name:      "Both Array Entries The Same",
			nums:      []int{1, 1},
			target:    2,
			expResult: []int{0, 1},
		},
		{
			name:      "Both Array Entries The Same Inverted Result",
			nums:      []int{1, 1},
			target:    2,
			expResult: []int{1, 0},
		},
		{
			name:      "Simple Array",
			nums:      []int{1, 2, 3, 4},
			target:    4,
			expResult: []int{0, 2},
		},
		{
			name:      "Simple Array",
			nums:      []int{1, 2, 3, 4},
			target:    4,
			expResult: []int{0, 2},
		},
	}
)

func TestTwoSum(t *testing.T) {
	for i, v := range testCases {
		result := twoSum(v.nums, v.target)
		if !checkExpectedResult(result, v.expResult) {
			t.Fatalf("(%v): %s: FAILED\n\nInput:%v, %v\nExpected: %v\nReceived: %v\n", i, v.name, v.nums, v.target, v.expResult, result)
		} else {
			t.Logf("(%v): %s: PASSED\n", i, v.name)
		}
	}
}

// Utility function to check in case of inversion

func checkExpectedResult(expResult []int, result []int) (isValid bool) {
	isValid = false
	if expResult[0] == result[0] {
		isValid = expResult[1] == result[1]
	} else if expResult[0] == result[1] {
		isValid = expResult[1] == result[0]
	}
	return
}
