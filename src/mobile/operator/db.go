package operator

import (
	"strconv"
)

type PhoneTree struct {
	root *PhoneTreeNode
}

type PhoneTreeNode struct {
	digit    string
	level    byte
	children map[byte]*PhoneTreeNode
	payload  *PhoneRangeNode
}

type PhoneRangeNode struct {
	right    [10]byte
	operator string
}

type PhoneTreeWalker struct {
	currentNode      *PhoneTreeNode
	currentPath      string
	nearestSplitNode *PhoneTreeNode
	visitedNodes     map[*PhoneTreeNode]interface{}
}

func NewTree(root *PhoneTreeNode) *PhoneTree {
	return &PhoneTree{root:root}
}

func NewTreeNode(digit string, level byte, payload *PhoneRangeNode) *PhoneTreeNode {
	children := make(map[byte]*PhoneTreeNode)
	return &PhoneTreeNode{digit:digit, level:level, children:children, payload:payload}
}


func (tree *PhoneTree) AddRange(fromStr string, toStr string, operator string) {
	var from, to [10]byte
	copy(from[:], fromStr)
	copy(to[:], toStr)
	node := tree.root
	for i := 1; i < 9; i++ {
		val := from[i]
		if _, ok := node.children[val]; ok {
			node = node.children[val]
		} else {
			child := NewTreeNode(string(val), byte(i)+1, nil)
			node.children[val] = child
			node = node.children[val]
		}
	}
	val := from[9]
	leaf := &PhoneRangeNode{to, operator}
	node.children[val] = NewTreeNode(string(val), 10, leaf)
}

func (tree *PhoneTree) AddPhone(phone string, operator string) {
	tree.AddRange(phone, phone, operator)
}

func (walker *PhoneTreeWalker) Valid() bool {
	if walker.currentNode.payload != nil {
		return true
	}

	for _, child := range walker.currentNode.children {
		if _, ok := walker.visitedNodes[child]; ok {
			continue
		}
		return true
	}

	return false
}

func (walker *PhoneTreeWalker) Walk(phone string) (string, bool) {
	if walker.currentNode.payload != nil {
		pathValue, _ := strconv.ParseInt(walker.currentPath, 10, 64)
		phoneValue, _ := strconv.ParseInt(phone, 10, 64)
		if phoneValue < pathValue {
			walker.currentNode = walker.nearestSplitNode
			walker.currentPath = walker.currentPath[0:walker.currentNode.level]

			return "", false
		}

		rightValue, _ := strconv.ParseInt(string(walker.currentNode.payload.right[:10]), 10, 64)
		if phoneValue > rightValue {
			walker.currentNode = walker.nearestSplitNode
			walker.currentPath = walker.currentPath[0:walker.currentNode.level]

			return "", false
		}

		op := walker.currentNode.payload.operator
		return op, true
	}

	var node *PhoneTreeNode

	val := []byte(phone[walker.currentNode.level:walker.currentNode.level+1])[0]

	if len(walker.currentNode.children) > 1 {
		walker.nearestSplitNode = walker.currentNode
	}
	candidate, ok := walker.currentNode.children[val]
	_, ok2 := walker.visitedNodes[candidate]
	if ok && !ok2 {
		node = walker.currentNode.children[val]
	} else {
		for key, child := range walker.currentNode.children {
			if _, ok := walker.visitedNodes[child]; ok {
				continue
			}
			if key < val || len(walker.currentNode.children) == 1 {
				node = child
			}
		}
	}
	walker.visitedNodes[walker.currentNode] = nil
	walker.currentNode = node
	walker.currentPath += node.digit

	return "", false
}

func (tree *PhoneTree) Find(phone string) string {
	walker := new(PhoneTreeWalker)
	walker.visitedNodes = make(map[*PhoneTreeNode]interface{})
	walker.currentNode = tree.root
	walker.currentPath = tree.root.digit
	for walker.Valid() {
		result, found := walker.Walk(phone)
		if found {
			return result
		}
	}

	return ""
}
