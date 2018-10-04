/*
Package main takes two numbers represented by linked lists and adds them together, as per LeetCode.com problem #2.

Written by Keith Drew, 2018.
*/
package main

import "fmt"

// ListNode : The struct used for linked lists.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := addNode(nil, createNode(3))
	head1 = addNode(head1, createNode(4))
	head1 = addNode(head1, createNode(2))

	printList(head1)

	head2 := addNode(nil, createNode(4))
	head2 = addNode(head2, createNode(6))
	head2 = addNode(head2, createNode(5))

	printList(head2)

	final := addTwoNumbers(head1, head2)

	printList(final)
}

func addNode(head *ListNode, newNode *ListNode) *ListNode {
	if head == nil {
		head = newNode
		newNode.Next = nil
	} else {
		newNode.Next = head
		head = newNode
	}

	return head
}

func printList(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head.Val)
	printList(head.Next)
}

func createNode(value int) *ListNode {
	returnValue := &ListNode{value, nil}
	return returnValue
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	// TODO: Hmm. Getting the lists in reverse means we can easily handle the carry, but what's the right recursion here?
	// Maybe recursion isn't the best answer?
	if l1.Val+l2.Val >= 10 {
		newNode := createNode((l1.Val + l2.Val) % 10)

	} else {
		newNode := createNode(l1.Val + l2.Val)
	}

	return nil
}
