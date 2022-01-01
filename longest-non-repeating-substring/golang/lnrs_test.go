package leetcode

import (
	"fmt"
	"testing"
)

type testData struct {
	name      string
	str       string
	expResult int
}

var (
	testCases []*testData = []*testData{
		{
			name:      "Single char string",
			str:       "a",
			expResult: 1,
		},
		{
			name:      "No characters repeat",
			str:       "abcdefghij",
			expResult: 10,
		},
		{
			name:      "One character repeats",
			str:       "abcdefghija",
			expResult: 10,
		},
		{
			name:      "Multiple repeating characters",
			str:       "abcabcbb",
			expResult: 3,
		},
		{
			name:      "All repeating characters",
			str:       "bbbb",
			expResult: 1,
		},
		{
			name:      "Find longest substring not subsection",
			str:       "pwwkew",
			expResult: 3,
		},
	}
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for i, v := range testCases {
		fmt.Printf("(%v) %s\n========================================\n", i, v.name)
		fmt.Printf("Input: %v\n", v.str)
		result := lengthOfLongestSubstring(v.str)
		if result != v.expResult {
			t.Fatalf("Result: FAILED\n\nExpected: %v\nReceived: %v\n", v.expResult, result)
		} else {
			fmt.Print("Result: PASSED\n\n")
		}
	}
}
