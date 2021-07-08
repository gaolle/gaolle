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

