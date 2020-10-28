---
title: 剑指offer
date: 2020-10-27 01:15:50
tags: 剑指offer
categories: 算法
---

# 剑指 Offer

#### [剑指 Offer 03. 数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)

```go
func findRepeatNumber(nums []int) int {
    tmp := make(map[int]bool)
    res := 0
    for _, val := range nums {
        if _, ok := tmp[val]; ok {
            tmp[val] = true
        } else {
            tmp[val] = false
        }

        if tmp[val] == true {
            res = val
            break
        }
    }
    return res
}
```

#### [剑指 Offer 04. 二维数组中的查找](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)

```go
func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    row := 0
    col := len(matrix[0]) - 1 // error len(matrix) - 1
    res := false

    for row < len(matrix) && col >= 0 {
        if matrix[row][col] < target {
            row += 1
        } else if matrix[row][col] > target {
            col -= 1
        } else if matrix[row][col] == target {
            res = true
            break
        }
    }
    return res
}
```

#### [剑指 Offer 05. 替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

```go
func replaceSpace(s string) string {
    res := ""
    for _, val := range s {
        if val == ' ' { 
            res += "%20"
        } else {
            res += string(val)
        }
    }
    return res
}
```

#### [剑指 Offer 06. 从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    var tmp []int

    for head != nil {
        tmp = append(tmp, head.Val)
        head = head.Next
    }
    var res []int
    for i := len(tmp) - 1; i >= 0; i-- {
        res = append(res, tmp[i])
    }
    return res
}
```

#### [剑指 Offer 07. 重建二叉树](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
		res := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:find(inorder, preorder[0])+1], inorder[:find(inorder, preorder[0])]),
		Right: buildTree(preorder[find(inorder, preorder[0])+1:], inorder[find(inorder, preorder[0])+1:]),
	}

	return res
}

func find(slice []int, tmp int) (index int) {
	for i, v := range slice {
		if v == tmp {
			index = i
			break
		}
	}
	return index
}
```

#### [剑指 Offer 10- I. 斐波那契数列](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)

```go
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	res := 0
	tmp0 := 0
	tmp1 := 1

	for i:= 2; i <= n; i++ {
		res = (tmp1 + tmp0) % 1000000007
		tmp0 = (tmp1) % 1000000007
		tmp1 = (res) % 1000000007
	}

	return res % 1000000007
}
```

#### [剑指 Offer 10- II. 青蛙跳台阶问题](https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/)

```go
func numWays(n int) int {
	if n == 0 || n ==1 {
		return 1
	}
	res := 0
	tmp0 := 1
	tmp1 := 1
	
	for i:= 2; i <=n; i++ {
		res = (tmp0 + tmp1)%1000000007    
		tmp0 = tmp1        
		tmp1 = res         
	}
    return res%1000000007
}
```

#### [剑指 Offer 11. 旋转数组的最小数字](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

```
func minArray(numbers []int) int {
	res := 0
	i := 0
	j := len(numbers) - 1
	for {
		if numbers[i] > numbers[j] {
			i++
		} else if numbers[i] < numbers[j] {
			j--
		} else if numbers[i] == numbers[j] {
			if i == j {
				res = numbers[i]
				break
			} else {
				i++
			}
		}
	}
	return res
}
```











