package leetcode

// This datastructure is given by LeetCode
type ListNode struct {
	Val  int
	Next *ListNode
}

// end

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = nil

	if l1 == nil || l2 == nil {
		return result
	}
	result = &ListNode{Val: l1.Val, Next: nil}
	rfinger := result
	addendFinger := l1.Next

	for addendFinger != nil {
		rfinger.Next = &ListNode{Val: addendFinger.Val, Next: nil}
		addendFinger = addendFinger.Next
		rfinger = rfinger.Next
	}
	addendFinger = l2
	// Set to a value BEFORE result so that in the loop we step in correctly
	rfinger = &ListNode{Val: 0, Next: result}
	carry := 0
	for addendFinger != nil {
		if rfinger.Next == nil {
			rfinger.Next = &ListNode{Val: 0, Next: nil}
		}
		rfinger = rfinger.Next
		rfinger.Val += carry
		rfinger.Val += addendFinger.Val
		if rfinger.Val > 9 {
			carry = 1
			rfinger.Val %= 10
		} else {
			carry = 0
		}
		addendFinger = addendFinger.Next
	}
	if carry != 0 {
		for carry != 0 {
			if rfinger.Next == nil {
				rfinger.Next = &ListNode{Val: carry, Next: nil}
				carry = 0
			} else {
				rfinger = rfinger.Next
				rfinger.Val += carry
				carry = 0
				if rfinger.Val > 9 {
					rfinger.Val -= 10
					carry = 1
				}
			}
		}
	}
	return result
}
