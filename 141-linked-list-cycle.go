package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	cache := []*ListNode{head}
	inArray := func(e *ListNode, ls []*ListNode) bool {
		for _, v := range ls {
			if e == v {
				return true
			}
		}
		return false
	}

	curr := head
	for curr.Next != nil {
		if inArray(curr.Next, cache) {
			return true
		} else {
			cache = append(cache, curr.Next)
		}
		curr = curr.Next
	}

	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if slow.Next == nil || fast.Next == nil {
			return false
		}
		if fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
