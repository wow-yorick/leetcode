package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	c := &ListNode{
		Val:  7,
		Next: nil,
	}
	b := &ListNode{
		Val:  6,
		Next: c,
	}
	a := &ListNode{
		Val:  5,
		Next: b,
	}

	e := &ListNode{
		Val:  4,
		Next: nil,
	}
	d := &ListNode{
		Val:  2,
		Next: e,
	}

	ret := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				l1: a,
				l2: d,
			},
			want: ret,
		},
	}
	for _, tt := range tests {
		fmt.Printf("a %v", a)
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
