package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		firstNum, secondNum, i, j int
		result                    *ListNode
	)
	if l1.Val == 0 {
		firstNum = 0
	} else {
		for {
			firstNum = firstNum + l1.Val*int(math.Pow10(i))
			i++
			if l1.Next == nil {
				break
			}
			l1 = l1.Next
		}
	}

	if l2.Val == 0 {
		secondNum = 0
	} else {
		for {
			secondNum = secondNum + l2.Val*int(math.Pow10(j))
			j++
			if l2.Next == nil {
				break
			}
			l2 = l2.Next
		}
	}
	sumNum := firstNum + secondNum
	memoryN := sumNum

	if sumNum == 0 {
		return linkInsert(&ListNode{Val: 0}, &ListNode{Val: 0})
	}

	//反转
	var newNum int
	for sumNum > 0 {
		ct := sumNum % 10
		newNum = newNum*10 + ct
		newNode := &ListNode{
			Val:  ct,
			Next: nil,
		}

		result = linkInsert(result, newNode)

		sumNum = sumNum / 10
	}
	fmt.Printf("first: %d second: %d the sum: %d reverse: %d", firstNum, secondNum, memoryN, newNum)
	return result

}

func linkInsert(link, node *ListNode) *ListNode {
	if link == nil {
		link = node
		return link
	}
	current := link
	for {
		if current.Next == nil {
			break
		}
		current = current.Next

	}
	current.Next = node
	return link
}
