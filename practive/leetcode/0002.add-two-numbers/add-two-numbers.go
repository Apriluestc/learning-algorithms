package problem0002

/*
[2] Add Two Numbers

https://leetcode.com/problems/add-two-numbers/description/

algorithms Medium (30.7%)
Total Accepted:    793.5K(793487)
Total Submissions: 2.6M(2582883)
Testcase Example:
`
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
`

You are given two **non-empty** linked lists representing two non-negative integers.
The digits are stored in **reverse order** and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	carry := 0
	p := l1
	q := l2
	for p != nil || q != nil {
		x := 0
		if p != nil {
			x = p.Val
		}
		y := 0
		if q != nil {
			y = q.Val
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return head.Next
}
