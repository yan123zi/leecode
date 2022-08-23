package main

import "fmt"

//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,2,1]
//输出：true
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 1475 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := []int{1, 2, 2, 1}
	nodeList := genNodeList(list)
	ret := isPalindrome(nodeList)
	fmt.Println(ret)
}

func genNodeList(valList []int) *ListNode {
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

//	自己的解题方式
func isPalindrome(head *ListNode) bool {
	curNode := head
	listSlice := []int{}
	for curNode.Next != nil {
		listSlice = append(listSlice, curNode.Val)
		curNode = curNode.Next
	}
	listSlice = append(listSlice, curNode.Val)
	if listSlice[0] != listSlice[len(listSlice)-1] {
		return false
	}
	right := len(listSlice) - 1
	for left, val := range listSlice {
		if left >= right {
			return true
		}
		if val == listSlice[right] {
			right--
		} else {
			return false
		}
	}
	return false
}

//	第二种方式，反转后半部分链表
//func ()  {
//
//}
