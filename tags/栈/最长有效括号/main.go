package main

import (
	"fmt"
)

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
//
// Related Topics 栈 字符串 动态规划 👍 1469 👎 0

func main() {
	str:="()()()()((()()())(())"
	longestValidParentheses(str)
}

func longestValidParentheses(s string) int {
	left,right:="(",")"
	matchMap:=map[int]int32{}
	storeSlice:=[]int32{}
	for index, ch := range s {
		matchMap[index]=ch
		if index==0 {
			continue
		}
		lastCh:=matchMap[index-1]
		if string(ch)==right&&string(lastCh)==left {
			storeSlice=append(storeSlice, lastCh,ch)
		}
	}
	fmt.Println(string(40),string(41))
	fmt.Println(storeSlice)
	//strconv.ParseInt()
	return 0
}