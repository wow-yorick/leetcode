package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	//var ret *ListNode
	len1 :=0
	len2 :=0
	for p1 != nil {
		len1++
		p1 = p1.Next
	}
	for p2 != nil {
		len2++
		p2 = p2.Next
	}

	p1 = headA
	p2 = headB
	if len1 > len2 {
		diff := len1 - len2
		for i :=0;i<diff;i++ {
			p1 = p1.Next
		}
	} else {
		diff := len2 - len1
		for i:=0; i<diff;i++ {
			p2 = p2.Next
		}
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
