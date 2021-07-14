```go
package main

import (
	"sort"
)

type Node struct {
	weight uint
	value  uint
	parent *Node
	left   *Node
	right  *Node
}

type Tree struct {
	root *Node
	leaf NodeList
	src  map[uint]uint
	code map[uint]string
}

type NodeList []*Node

func (nl NodeList) len() int {
	return len(nl)
}

func (nl NodeList) swap(i, j int) {
	nl[i], nl[j] = nl[j], nl[i]
}

func (nl NodeList) less(i, j int) bool {
	return nl[i].weight < nl[j].weight
}

func (nl NodeList) sort() {
	sort.Slice(nl, func(i, j int) bool {
		if nl.less(i, j) {
			return true
		}
		return false
	})
}

func newTree(src map[uint]uint) *Tree {
	tree := &Tree{
		src: src,
	}
	tree.init()
	tree.build()
	tree.parse()
	return tree
}

func (t *Tree) init() {
	if len(t.src) <= 1 {
		panic("src len not enough")
	}

	t.leaf = make(NodeList, len(t.src))
	var i int
	for value, weight := range t.src {
		node := &Node{
			weight: weight,
			value:  value,
		}
		t.leaf[i] = node
		i += 1
	}
	t.leaf.sort()
}

func (t *Tree) build() {
	nodeList := t.leaf
	for nodeList.len() > 1 {
		temp := &Node{
			weight: nodeList[0].weight + nodeList[1].weight,
			left:   nodeList[0],
			right:  nodeList[1],
		}
		nodeList[0].parent = temp
		nodeList[1].parent = temp
		// 将新生成的节点加入，根据权重重新排序
		nodeList = nodeList[2:]
		nodeList = append(nodeList, temp)
		nodeList.sort()
	}
	t.root = nodeList[0]
}

func (t *Tree) parse() {
	if t.root == nil {
		return
	}
	t.code = make(map[uint]string)
	preOrder(t.code, t.root, "", "")
}

// 前序排列生成编码
func preOrder(m map[uint]string, n *Node, s, arr string) {
	if n == nil {
		return
	}

	arr += s
	if n.left == nil {
		m[n.value] = arr
		return
	} else {
		preOrder(m, n.left, "0", arr)
	}

	if n.right == nil {
		m[n.value] = arr
		return
	} else {
		preOrder(m, n.right, "1", arr)
	}
}

func main() {
	src := map[uint]uint{
		101: 5,
		102: 4,
		103: 3,
		104: 2,
		105: 1,
	}
	tree := newTree(src)
	println("%s", tree)
}

```





```rust
use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::Rc;

#[derive(Debug, Eq, Ord, PartialEq, PartialOrd, Clone)]
struct HuffmanNode {
    weight: isize,
    val: isize,
    parent: Option<Rc<RefCell<HuffmanNode>>>,
    left: Option<Rc<RefCell<HuffmanNode>>>,
    right: Option<Rc<RefCell<HuffmanNode>>>,
}

#[derive(Debug)]
struct HuffmanTree {
    root: Option<Rc<RefCell<HuffmanNode>>>,
    leaf: Vec<Option<Rc<RefCell<HuffmanNode>>>>,
    src: HashMap<isize, isize>,
    code: HashMap<isize, String>,
}

impl HuffmanTree {
    fn new(src: HashMap<isize, isize>) -> HuffmanTree {
        if src.len() <= 1 {
            panic!("src len not enough");
        }

        let mut tree = HuffmanTree {
            root: None,
            leaf: Vec::with_capacity(src.len()),
            src: HashMap::new(),
            code: HashMap::new(),
        };

        for (value, weight) in src.iter() {
            let node = Some(Rc::new(RefCell::new(HuffmanNode {
                weight: *weight,
                val: *value,
                parent: None,
                left: None,
                right: None,
            })));

            tree.leaf.push(node);
        }
        tree.src = src;
        sort_by_weight(&mut tree.leaf);

        tree
    }

    fn build(&mut self) {
        let mut leaf = self.leaf.clone();
        while leaf.len() > 1 {
            let temp = Some(Rc::new(RefCell::new(
                HuffmanNode {
                    weight: leaf[0].as_ref().unwrap().borrow_mut().weight + leaf[1].as_ref().unwrap().borrow_mut().weight,
                    val: 0,
                    parent: None,
                    left: leaf[0].clone(),
                    right: leaf[1].clone(),
                }
            )));
            leaf[0].as_ref().unwrap().borrow_mut().parent = temp.clone();
            leaf[1].as_ref().unwrap().borrow_mut().parent = temp.clone();
            leaf = leaf[2..].to_vec();
            leaf.push(temp);
            sort_by_weight(&mut leaf);
        }
        self.root = leaf[0].clone();
    }

    fn parse(&mut self) {
        if self.root.is_none() {
            return;
        }
        pre_order(&mut self.code, &self.root, String::new(), String::new());
    }
}

fn sort_by_weight(leaf: &mut Vec<Option<Rc<RefCell<HuffmanNode>>>>) {
    leaf.sort_by(|x, y|
        x.as_ref().unwrap().borrow_mut().weight.cmp(&y.as_ref().unwrap().borrow_mut().weight));
}

fn pre_order(m: &mut HashMap<isize, String>, node: &Option<Rc<RefCell<HuffmanNode>>>, s: String, mut arr: String) {
    if node.is_none() {
        return;
    }

    arr.push_str(&s);
    if node.as_ref().unwrap().borrow_mut().left.is_none() {
        m.insert(node.as_ref().unwrap().borrow_mut().val, arr.clone());
        return;
    } else {
        pre_order(m, &node.as_ref().unwrap().borrow_mut().left, String::from("0"), arr.clone());
    }
    if node.as_ref().unwrap().borrow_mut().right.is_none() {
        m.insert(node.as_ref().unwrap().borrow_mut().val, arr.clone());
        return;
    } else {
        pre_order(m, &node.as_ref().unwrap().borrow_mut().right, String::from("1"), arr.clone());
    }
}

fn main() {
    let src: HashMap<isize, isize> = [(101, 5), (102, 4), (103, 3), (104, 2), (105, 1)].iter().cloned().collect();
    // println!("{:#?}", src);

    let mut tree = HuffmanTree::new(src);
    tree.build();
    tree.parse();
    println!("{:?}", tree.code);
    // println!("{:?}", tree);
}
```

