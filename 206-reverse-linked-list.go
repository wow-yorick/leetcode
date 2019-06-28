package main

import (
	"encoding/json"
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head

	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}

	return prev
}

func main() {

	in := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
				Next:nil,
			},
		},
	}

	out := reverseList(in)

	s,_ := json.Marshal(in)
	d,_ := json.Marshal(out)

	fmt.Printf("in: %s  out: %s", s, d)
}
