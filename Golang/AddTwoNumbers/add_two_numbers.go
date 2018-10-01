/* Written by Keith Drew, 2018
* This is a solution to LeetCode Problem #2: "Add Two Numbers," written in Go
 */
package main

import "fmt"

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

	// TODO: implement the solution :D
	
	return nil
}
