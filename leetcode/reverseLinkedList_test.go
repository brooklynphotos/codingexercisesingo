package leetcode

import "testing"

func TestReverseLinkedList(t *testing.T) {
	llist := makeList(1, 2, 3)
	reversedL := reverseList(llist)
	expected := 3
	if reversedL.Val != expected {
		t.Errorf("Expected list[0] %d but instead was %d", expected, reversedL.Val)
	}
	expectedLength := 3
	if reversedL.length() != expectedLength {
		t.Errorf("Expected size = %d but instead was %d", expectedLength, reversedL.length())
	}

}
func TestReverseLinkedListSingle(t *testing.T) {
	llist := makeList(1)
	reversedL := reverseList(llist)
	expected := 1
	if reversedL.Val != expected {
		t.Errorf("Expected list[0] %d but instead was %d", expected, reversedL.Val)
	}
	expectedLength := 1
	if reversedL.length() != expectedLength {
		t.Errorf("Expected size = %d but instead was %d", expectedLength, reversedL.length())
	}

}
func TestReverseLinkedEmpty(t *testing.T) {
	var llist *ListNode
	reversedL := reverseList(llist)
	if reversedL != nil {
		t.Errorf("Expected nothing but got %d", reversedL.length())
	}

}

func makeList(first int, ints ...int) *ListNode {
	llist := ListNode{first, nil}
	node := &llist
	for _, v := range ints {
		newNode := &ListNode{v, nil}
		node.Next = newNode
		node = newNode
	}
	return &llist
}

func (node *ListNode) length() int {
	count := 0
	for n := node; n != nil; n = n.Next {
		count++
	}
	return count
}
