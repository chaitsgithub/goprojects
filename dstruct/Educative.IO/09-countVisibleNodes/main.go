package main

import (
	"fmt"
)

type TreeNode struct {
	data                int
	leftNode, rightNode *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) InsertNode(num int) {
	newNode := &TreeNode{data: num}

	if t.root == nil {
		t.root = newNode
	} else {
		t.addNode(newNode, t.root)
	}
}

func (t *Tree) addNode(newNode *TreeNode, currNode *TreeNode) {

	if currNode.data > newNode.data {
		if currNode.leftNode == nil {
			currNode.leftNode = newNode
			return
		} else {
			t.addNode(newNode, currNode.leftNode)
		}
	} else {
		if currNode.rightNode == nil {
			currNode.rightNode = newNode
			return
		} else {
			t.addNode(newNode, currNode.rightNode)
		}
	}
}

func (t *Tree) printPreOrder(node *TreeNode, ind string) {
	if node == nil {
		return
	}
	fmt.Printf("%s : %d\n", ind, node.data)
	t.printPreOrder(node.leftNode, "LEFT")
	t.printPreOrder(node.rightNode, "RIGHT")
}

func (t *Tree) getMaxTreeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	} else {
		return getMax(t.getMaxTreeDepth(node.leftNode), t.getMaxTreeDepth(node.rightNode)) + 1
	}
	return -1
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	t := &Tree{root: nil}
	arr := []int{50, 40, 70, 60, 20, 99, 45}

	for _, i := range arr {
		t.InsertNode(i)
	}
	//t.printPreOrder(t.root, "ROOT")

	fmt.Printf("Max Tree Depth is %d", t.getMaxTreeDepth(t.root))
}
