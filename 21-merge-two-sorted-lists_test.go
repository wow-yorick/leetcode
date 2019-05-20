package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func listAppendEle(list *ListNode, ele int) {
	node := &ListNode{
		Val: ele,
	}
	listAppendChild(list, node)
}

func listAppendChild(list, node *ListNode) {
	current := list
	for {
		if current.Val == 0 && current.Next == nil {
			list.Val = node.Val
			list.Next = node.Next
			break
		}
		if current.Next == nil {
			current.Next = node
			break
		}
		current = current.Next
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	list1 := &ListNode{}
	listAppendEle(list1, 1)
	listAppendEle(list1, 2)
	listAppendEle(list1, 3)
	j, _ := json.Marshal(list1)
	fmt.Printf("LIST1 %v \n", string(j))

	list2 := &ListNode{}
	listAppendEle(list2, 4)
	listAppendEle(list2, 5)
	listAppendEle(list2, 6)
	k, _ := json.Marshal(list2)
	fmt.Printf("LIST2 %v \n", string(k))

	list3 := &ListNode{}
	listAppendEle(list3, 1)
	listAppendEle(list3, 2)
	listAppendEle(list3, 3)
	listAppendEle(list3, 4)
	listAppendEle(list3, 5)
	listAppendEle(list3, 6)

	m, _ := json.Marshal(list3)
	fmt.Printf("LIST3 %v \n", string(m))

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				l1: list1,
				l2: list2,
			},
			want: list3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoListsRight(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
