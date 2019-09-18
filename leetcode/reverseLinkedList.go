/**
 * https://leetcode.com/problems/reverse-linked-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return reverseListFor(head.Next, &ListNode{head.Val, nil})
}

func reverseListFor(head *ListNode, cumulator *ListNode) *ListNode {
	if head == nil {
		return cumulator
	}
	return reverseListFor(head.Next, &ListNode{head.Val, cumulator})
}

type ListNode struct {
	Val  int
	Next *ListNode
}
