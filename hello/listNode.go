package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func constructList(array []int) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for _, value := range array {
		tmp := &ListNode{}
		tmp.Val = value
		head.Next = tmp
		head = head.Next
	}
	head.Next = nil
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" ")
		head = head.Next
	}
	fmt.Println()
}
