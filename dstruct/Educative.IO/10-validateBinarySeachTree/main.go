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

func (t *Tree) isValidBinarySearchTree(node *TreeNode, maxSoFar, minSoFar int) bool {
	if node == nil {
		return true
	}
	if node.data < minSoFar {
		minSoFar = node.data
	}

	if node.data > maxSoFar {
		maxSoFar = node.data
	}

	if node.leftNode.data > minSoFar || node.rightNode.data < maxSoFar {
		fmt.Printf("Invalid Node; N : %d, L: %d, R: %d, min: %d, max: %d \n", node.data, node.leftNode.data, node.rightNode.data, minSoFar, maxSoFar)
		return false
	}

	return t.isValidBinarySearchTree(node.leftNode, maxSoFar, minSoFar) && t.isValidBinarySearchTree(node.leftNode, maxSoFar, minSoFar)
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

	fmt.Printf("Binary Search Tree is %v", t.isValidBinarySearchTree(t.root, t.root.data, t.root.data))
}
