package main

import "fmt"

func main() {
	list1 := genNodeList([]int{1,2,3})
	list2 := genNodeList([]int{1,3,4})
	res := mergeTwoLists1(list1, list2)
	fmt.Println(res)
}

func genNodeList(vals []int) *ListNode {
	var head *ListNode
	curNode := head
	for _, val := range vals {
		node := &ListNode{Val: val}
		if curNode == nil {
			curNode = node
			head = curNode
		} else {
			curNode.Next = node
			curNode = curNode.Next
		}
	}
	return head
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next=mergeTwoLists1(list1.Next,list2)
		return list1
	}else{
		list2.Next=mergeTwoLists1(list1,list2.Next)
		return list2
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list2 == nil {
		return list1
	} else if list1 == nil {
		return list2
	}
	resNode := list2
	for list1 != nil {
		addNode := &ListNode{Val: list1.Val, Next: nil}
		var preNode *ListNode
		var nextNode *ListNode
		for list2 != nil {
			preNode = list2
			nextNode = list2.Next
			if addNode.Val <= preNode.Val {
				addNode.Next = preNode
				list2 = addNode
				resNode = list2
				break
			}
			if nextNode != nil {
				if addNode.Val <= nextNode.Val {
					preNode.Next = addNode
					addNode.Next = nextNode
					break
				}
			} else {
				preNode.Next = addNode
				break
			}

			list2 = list2.Next
		}
		list1 = list1.Next
	}
	return resNode
}
