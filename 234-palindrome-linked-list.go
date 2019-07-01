package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var(
		slow  = head
		fast = head
	)
	for fast.Next != nil &&  fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	last := slow.Next
	pre := head

	for last.Next != nil {
		tmp := last.Next
		last.Next = tmp.Next
		tmp.Next = slow.Next
		slow.Next = tmp
	}
	for slow.Next != nil {
		slow = slow.Next
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
	}
	return true

}