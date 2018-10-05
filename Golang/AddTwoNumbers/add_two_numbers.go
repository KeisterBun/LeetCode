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

// Our main routine
func main() {
	// Set up and print some lists, for demo/sanity checks
	head1 := addNode(nil, &ListNode{8, nil})
	head1 = addNode(head1, &ListNode{1, nil})

	printList(head1)

	head2 := addNode(nil, &ListNode{0, nil})

	printList(head2)

	final := addTwoNumbers(head1, head2)

	printList(final)
}

// addNode adds a new node to the front of the list provided, returning a pointer to the new head
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

// printList prints out the given list in a readable format
func printList(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Printf("%d ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Printf("\n")
}

// addTwoNumbers is the LeetCode ready solution to problem #2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Make sure we're sane here
	if l1 == nil || l2 == nil {
		return nil
	}

	// Set up first operation outside of loop
	var carry int
	if l1.Val+l2.Val >= 10 {
		carry = 1
	}
	tmpNode := &ListNode{((l1.Val + l2.Val) % 10), nil}
	l1 = l1.Next
	l2 = l2.Next
	newListHead := tmpNode

	// Loop over all list elements until the list length differs
	for l1 != nil && l2 != nil {
		// Add our new node to the end of the list and move the tmp pointer forward in the list
		tmpNode.Next = &ListNode{((l1.Val + l2.Val + carry) % 10), nil}
		tmpNode = tmpNode.Next

		// Set the carry value
		if l1.Val+l2.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		// Advance the lists
		l1 = l1.Next
		l2 = l2.Next
	}

	// This is the same loop as above for the case where l1 list continues
	for l1 != nil {
		tmpNode.Next = &ListNode{(l1.Val + carry) % 10, nil}
		tmpNode = tmpNode.Next

		if l1.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		l1 = l1.Next
	}

	// This is the same loop as above above for the case where l2 list continues
	for l2 != nil {
		tmpNode.Next = &ListNode{(l2.Val + carry) % 10, nil}
		tmpNode = tmpNode.Next

		if l2.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		l2 = l2.Next
	}

	// Handle any final carry
	if carry == 1 {
		tmpNode.Next = &ListNode{1, nil}
	}

	// Return the new list
	return newListHead
}
