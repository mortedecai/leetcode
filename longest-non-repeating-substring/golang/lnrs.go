package leetcode

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var subStr string
	maxLength := 1

	firstIdx := 0

	for i, v := range s {
		subStr = s[firstIdx:i]
		if loc := strings.IndexRune(subStr, v); loc >= 0 {
			s2 := s[firstIdx:i]
			if l := len(s2); l > maxLength {
				maxLength = l
			}
			firstIdx += loc + 1
		}
	}
	subStr = s[firstIdx:]
	if l := len(subStr); l > maxLength {
		maxLength = l
	}

	return maxLength
}
