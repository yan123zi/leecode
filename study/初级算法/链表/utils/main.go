package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenNodeList(valList []int) *ListNode {
	var head *ListNode
	curNode := head
	for _, val := range valList {
		if curNode == nil {
			curNode = &ListNode{Val: val, Next: nil}
			head = curNode
			continue
		}
		node := &ListNode{Val: val, Next: nil}
		curNode.Next = node
		curNode = curNode.Next
	}

	return head
}