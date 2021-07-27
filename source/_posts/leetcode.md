## 剑指Offer

#### [剑指 Offer 03. 数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)

```rust
use std::collections::HashMap;
impl Solution {
    pub fn find_repeat_number(nums: Vec<i32>) -> i32 {
        let mut m: HashMap<i32, i32> = HashMap::new();
        for n in nums {
            let cnt = m.entry(n).or_insert(0);
            *cnt += 1;
            if *cnt > 1 {
                return n;
            }
        }
        0
    }
}
```

#### [剑指 Offer 04. 二维数组中的查找](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)

```rust
impl Solution {
    pub fn find_number_in2_d_array(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        if matrix.is_empty() {
            return false;
        }

        let x = matrix.len();
        let y = matrix[0].len();

        for i in 0..x {
            for j in (0..y).rev() {
                if matrix[i][j] == target {
                    return true;
                } else if matrix[i][j] < target {
                    break;
                } else if matrix[i][j] > target {
                    continue;
                }
            }
        }
        false
    }
}
```

#### [剑指 Offer 05. 替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

```rust
impl Solution {
    pub fn replace_space(s: String) -> String {
        let mut r = String::new();
        for c in s.chars() {
            if c == ' ' {
                r.push_str("%20");
            } else {
                r.push(c);
            }
        }
        r
    }
}
```

#### [剑指 Offer 06. 从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

```rust
// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn reverse_print(mut head: Option<Box<ListNode>>) -> Vec<i32> {
        let mut r = Vec::new();
        while head.is_some() {
            r.push(head.clone().unwrap().val);
            head = head.unwrap().next;
        }
        r.reverse();
        r
    }
}
```

#### [剑指 Offer 07. 重建二叉树](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)

```rust
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        if preorder.is_empty() {
            return None;
        }

        let r = Some(Rc::new(RefCell::new(TreeNode {
            val: preorder[0],
            left: None,
            right: None,
        })));
        let index = Solution::search_index(&inorder, preorder[0]);
        r.as_ref().unwrap().borrow_mut().left = Solution::build_tree(preorder[1..=index].to_vec(), inorder[..index].to_vec());
        r.as_ref().unwrap().borrow_mut().right = Solution::build_tree(preorder[index + 1..].to_vec(), inorder[index + 1..].to_vec());
        r
    }


    pub fn search_index(numbers: &Vec<i32>, target: i32) -> usize {
        for i in 0..numbers.len() {
            if numbers[i] == target {
                return i;
            }
        }
        0
    }
}
```



#### [剑指 Offer 09. 用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

```rust
struct CQueue {
    list: Vec<i32>
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl CQueue {
    fn new() -> Self {
        CQueue {
            list: Vec::new(),
        }
    }

    fn append_tail(&mut self, value: i32) {
        self.list.push(value);
    }

    fn delete_head(&mut self) -> i32 {
        return if self.list.is_empty() {
            -1
        } else {
            let i = self.list[0];
            self.list = self.list[1..self.list.len()].to_vec();
            i
        }
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * let obj = CQueue::new();
 * obj.append_tail(value);
 * let ret_2: i32 = obj.delete_head();
 */
```



#### [剑指 Offer 10- I. 斐波那契数列](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)

```rust
impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }
        if n == 1 {
            return 1;
        }

        let mut a = 0;
        let mut b = 1;
        let mut c = 1;
        for _i in 2..n {
            a = b % 1000000007;
            b = c % 1000000007;
            c = (a + b) % 1000000007;
        }
        c
    }
}
```

#### [剑指 Offer 10- II. 青蛙跳台阶问题](https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/)

```rust
impl Solution {
    pub fn num_ways(n: i32) -> i32 {
        if n == 0 {
            return 1;
        }
        if n == 1 {
            return 1;
        }

        let mut a = 1;
        let mut b = 1;
        let mut c = 2;

        for i in 2..n {
            a = b % 1000000007;
            b = c % 1000000007;
            c = (a + b) % 1000000007;
        }
        c
    }
}
```

#### [剑指 Offer 11. 旋转数组的最小数字](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

```rust
impl Solution {
    pub fn min_array(numbers: Vec<i32>) -> i32 {
        let mut i = 0;
        let mut j = numbers.len()-1;
        loop {
            if numbers[i] > numbers[j] {
                i += 1;
            } else if numbers[i] < numbers[j] {
                j -= 1;
            } else if numbers[i] == numbers[j] {
                if i == j {
                    return numbers[i];
                } else {
                    i += 1;
                }
            }
        }
    }
}
```

#### [剑指 Offer 17. 打印从1到最大的n位数](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

```
impl Solution {
    pub fn print_numbers(n: i32) -> Vec<i32> {
        let mut max = 1;
        for i in 0..n {
            max *= 10;
        }
        let mut res = vec![];
        for i in 1..max {
            res.push(i);
        }
        res
    }
}
```



#### [剑指 Offer 21. 调整数组顺序使奇数位于偶数前面](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)

```
impl Solution {
    pub fn exchange(nums: Vec<i32>) -> Vec<i32> {
        let mut res = Vec::new();
        for n in nums {
            if n % 2 == 0 {
                res.push(n);
            } else {
                res.insert(0, n);
            }
        }
        res
    }
}
```



#### [剑指 Offer 27. 二叉树的镜像](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)

```rust
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn mirror_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        match root {
            None => root,
            Some(node) => {
                let l = node.as_ref().borrow_mut().left.take();
                let r = node.as_ref().borrow_mut().right.take();
                node.as_ref().borrow_mut().left = r;
                node.as_ref().borrow_mut().right = l;
                Solution::mirror_tree(node.as_ref().borrow_mut().left.clone());
                Solution::mirror_tree(node.as_ref().borrow_mut().right.clone());
                Some(node)
            }
        }
    }
}
```



#### [剑指 Offer 28. 对称的二叉树](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/)

```rust
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn is_symmetric(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root.is_none() {
            return true;
        }
        let r = root.as_ref().unwrap().borrow_mut();
        Solution::symmetric(r.left.clone(), r.right.clone())
    }

    pub fn symmetric(l: Option<Rc<RefCell<TreeNode>>>, r: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if l.is_none() && r.is_none() {
            return true;
        }
        if l.is_none() || r.is_none() {
            return false;
        }
        let lv = l.as_ref().unwrap().borrow_mut();
        let rv = r.as_ref().unwrap().borrow_mut();
        if lv.val != rv.val {
            return false;
        }
        Solution::symmetric(lv.left.clone(), rv.right.clone()) && Solution::symmetric(lv.right.clone(), rv.left.clone())
    }
}
```



#### [剑指 Offer 29. 顺时针打印矩阵](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)

```rust
impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        if matrix.is_empty() {
            return vec![];
        }
        let (mut x_min, mut x_max, mut y_min, mut y_max) = (0 as i32, (matrix.len() - 1) as i32, 0 as i32, (matrix[0].len() - 1) as i32);
        let (mut x, mut y) = (0, 0);
        let mut res = vec![];
        loop {
            if y <= y_max as usize {
                for yy in y..=y_max as usize {
                    res.push(matrix[x][yy as usize]);
                }
                y = y_max as usize;
                x_min += 1;
                x += 1;
            }
            if x_min > x_max {
                return res;
            }
            if x <= x_max as usize {
                for xx in x..=x_max as usize {
                    res.push(matrix[xx][y]);
                }
                x = x_max as usize;
                y_max -= 1;
                if y != 0 {
                    y -= 1;
                }
            }
            if y_min > y_max {
                return res;
            }
            if y >= y_min as usize {
                for yy in (y_min as usize..=y).rev() {
                    res.push(matrix[x][yy]);
                }
                y = y_min as usize;
                x_max -= 1;
                if x != 0 {
                    x -= 1;
                }
            }
            if x_min > x_max {
                return res;
            }
            if x >= x_min as usize {
                for xx in (x_min as usize..=x).rev() {
                    res.push(matrix[xx][y]);
                }
                x = x_min as usize;
                y_min += 1;
                y += 1;
            }
            if y_min > y_max {
                return res;
            }
        }
    }
}
```

#### [剑指 Offer 22. 链表中倒数第k个节点](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)

```rust
// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn get_kth_from_end(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        let mut len = 0;
        let mut h = head.clone();
        while h.is_some() {
            len += 1;
            h = h.unwrap().next.clone();
        }
        let mut t = head.clone();
        while len - k > 0 {
            len -= 1;
            t = t.unwrap().next.clone();
        }
        t
    }
}
```

#### [剑指 Offer 47. 礼物的最大价值](https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/)

```rust
use std::cmp::max;
impl Solution {
    pub fn max_value(grid: Vec<Vec<i32>>) -> i32 {
        let (x_max, y_max) = (grid.len(), grid[0].len());
        let mut dp = grid.clone();

        for y in 1..y_max {
            dp[0][y] = grid[0][y] + dp[0][y - 1];
        }
        for x in 1..x_max {
            dp[x][0] = grid[x][0] + dp[x - 1][0];
        }
        for x in 1..x_max {
            for y in 1..y_max {
                dp[x][y] = grid[x][y] + max(dp[x - 1][y], dp[x][y - 1]);
            }
        }
        dp[x_max - 1][y_max - 1]
    }
}
```

