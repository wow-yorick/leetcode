package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummyHead = &ListNode{Val: 0}
		p         = l1
		q         = l2
		curr      = dummyHead
		carry     = 0
		x, y      int
	)

	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}
	return dummyHead.Next

}
