## [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

### 迭代法

![image-20210520000408648](image/image-20210520000408648.png)

核心在于辅助空节点的使用

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        temp := head.Next
        head.Next = prev
        prev = head
        head = temp
    }
    return prev
}
```

### 递归法

每当递归函数调用自身时，它都会将给定的问题拆解为**子问题**。递归调用继续进行，直到到子问题成为一个不可以拆分的、可以直接求解的最简单问题。

递归的秘诀就是，**千万不要跳进去**，人脑不是电脑，压 3 个栈估计就懵了。

而是要根据**递归函数的定义**，认为它可以帮你处理好一切中间过程，你要做的就是收尾工作。

![image-20210526235444011](image/image-20210526235444011.png)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    last := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return last
}
```

## [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

不要被两个链表干扰，本质还是按照一个节点来处理

## [24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

注意本题的最小子问题是两个节点要看为一个整体

## 参考

https://labuladong.gitee.io/algo/2/17/16/

