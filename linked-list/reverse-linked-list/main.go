// Reverse Linked List. Solution in Golang language
// Given the head of a singly linked list, reverse the list, and return the reversed list.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev ListNode

	for curr.Next != nil {
		nxt := curr.Next
		curr.Next = &prev
		prev = *curr
		curr = nxt
	}
	return &prev
}

func main() {
	var in ListNode = ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{}}}}}}
	fmt.Println(reverseList(&in))
}
