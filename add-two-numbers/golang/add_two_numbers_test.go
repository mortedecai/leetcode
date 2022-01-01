package leetcode

import (
	"fmt"
	"testing"
)

type testData struct {
	name      string
	num1      string
	num2      string
	expResult string
}

var (
	testCases []*testData = []*testData{
		{
			name:      "Zero Plus Zero",
			num1:      "0",
			num2:      "0",
			expResult: "0",
		},
		{
			name:      "1 + 1",
			num1:      "1",
			num2:      "1",
			expResult: "2",
		},
		{
			name:      "1 + 9",
			num1:      "1",
			num2:      "9",
			expResult: "10",
		},
		{
			name:      "10 + 2",
			num1:      "10",
			num2:      "2",
			expResult: "12",
		},
		{
			name:      "1 + 10",
			num1:      "1",
			num2:      "10",
			expResult: "11",
		},
		{
			name:      "342 + 465",
			num1:      "342",
			num2:      "465",
			expResult: "807",
		},
		{
			name:      "9999999 + 9999",
			num1:      "9999999",
			num2:      "9999",
			expResult: "10009998",
		},
	}
)

func TestAddTwoNumbers(t *testing.T) {
	for i, v := range testCases {
		n1 := convertFromString(v.num1)
		n2 := convertFromString(v.num2)
		result := addTwoNumbers(n1, n2)
		expResult := convertFromString(v.expResult)
		if !checkExpectedResult(result, expResult) {
			t.Fatalf("(%v): %s: FAILED\n\nInput: %v, %v\nExpected: %v\nReceived: %v\n", i, v.name, v.num1, v.num2, stringRep(expResult), stringRep(result))
		} else {
			t.Logf("(%v): %s: PASSED\n", i, v.name)
		}
	}
}

// Utility function to return the string representation of a ListNode list
func stringRep(v *ListNode) string {
	if nil == v {
		return "<nil>"
	}
	finger := v
	result := ""
	hasValue := false

	for finger != nil {
		if hasValue {
			result = fmt.Sprintf("%s -> ", result)
		}
		result = fmt.Sprintf("%s(%v)", result, finger.Val)
		finger = finger.Next
	}
	return result
}

// Utility function to compare to linked lists
func checkExpectedResult(result *ListNode, expResult *ListNode) bool {
	f1 := result
	f2 := expResult

	for f1 != nil {
		if f2 == nil || f1.Val != f2.Val {
			return false
		}
		f1 = f1.Next
		f2 = f2.Next
	}

	return f1 == f2
}

// Utility function to convert string to a ListNode pointer

func convertFromString(val string) *ListNode {
	var finger *ListNode = nil
	for _, c := range val {
		v := int(c - '0')
		node := &ListNode{Val: v, Next: finger}
		finger = node
	}

	return finger
}
