package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	listTraverse := func(l1 *ListNode) []int {

		fmt.Println("int", l1)
		ret := []int{}
		if l1 == nil {
			return ret
		}
		current := l1
		for {
			ret = append(ret, current.Val)
			if current.Next == nil {
				break
			}
			current = current.Next
		}
		return ret
	}

	f := listTraverse(l1)
	l := listTraverse(l2)
	f = append(f, l...)
	sort.Ints(f)

	dumpHeader := &ListNode{
		Val: 0,
	}
	currentList := dumpHeader
	for _, v := range f {
		currentList.Next = &ListNode{
			Val: v,
		}
		currentList = currentList.Next
	}
	return dumpHeader.Next
}

func mergeTwoListsRight(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRight(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsRight(l2.Next, l1)
		return l2
	}
}
