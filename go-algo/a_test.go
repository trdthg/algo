package goalgo

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	set := make(map[*ListNode]struct{})
	for head.Next != nil {
		if _, ok := set[head]; ok {
			return true
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return false
}

func TestC(t *testing.T) {
	node := ListNode{Val: 3, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 0, Next: nil}
	node4 := ListNode{Val: -4, Next: nil}
	node.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2
	if res := hasCycle(&node); res == false {
		t.Errorf("res shouldn't not be %t", res)
	}

}
