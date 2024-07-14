package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	key   int
	value interface{}
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (bst *BinarySearchTree) Insert(key int, value interface{}) {
	bst.root = bst.insertNode(bst.root, key, value)
}

func (bst *BinarySearchTree) Find(key int) bool {
	currentNode := bst.root
	for currentNode != nil {
		if key == currentNode.key {
			return true
		}
		if key < currentNode.key {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return false
}

func (bst *BinarySearchTree) Remove(key int) {
	bst.root = bst.removeNode(bst.root, key)
}

func (bst *BinarySearchTree) TraverseInOrder() string {
	if bst.root == nil {
		return ""
	}
	var nodesString string
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node.left != nil {
			traverse(node.left)
		}
		nodesString += fmt.Sprintf(" %d", node.key)
		if node.right != nil {
			traverse(node.right)
		}
	}
	traverse(bst.root)
	return nodesString[1:]
}

func (bst *BinarySearchTree) insertNode(node *TreeNode, key int, value interface{}) *TreeNode {
	if node == nil {
		return &TreeNode{key: key, value: value}
	}
	if key < node.key {
		node.left = bst.insertNode(node.left, key, value)
	} else if key > node.key {
		node.right = bst.insertNode(node.right, key, value)
	} else {
		panic(errors.New("cannot insert a node with a duplicate key"))
	}
	return node
}

func (bst *BinarySearchTree) findMin(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BinarySearchTree) removeMin(node *TreeNode) *TreeNode {
	if node.left == nil {
		return node.right
	}
	node.left = bst.removeMin(node.left)
	return node
}

func (bst *BinarySearchTree) removeNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = bst.removeNode(node.left, key)
	} else if key > node.key {
		node.right = bst.removeNode(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}

		target := node
		successor := bst.findMin(target.right)
		successor.right = bst.removeMin(target.right)
		successor.left = target.left
		return successor
	}
	return node
}
