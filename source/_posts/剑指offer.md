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































