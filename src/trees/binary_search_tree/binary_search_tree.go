package binarysearchtree

import (
	"errors"
)

type node struct {
	key    int
	value  interface{}
	left   *node
	right  *node
	parent *node
}

func (n *node) insert(key int, value interface{}) bool {
	if key == n.key {
		n.value = value
		return false

	} else if key < n.key {
		if n.left == nil {
			n.left = &node{key: key, value: value, parent: n}
			return true
		}
		return n.left.insert(key, value)

	} else {
		if n.right == nil {
			n.right = &node{key: key, value: value, parent: n}
			return true
		}
		return n.right.insert(key, value)
	}
}

func (n *node) delete(key int) (interface{}, bool, bool) {
	if key == n.key {
		var value interface{} = n.value
		if n.left == nil && n.right == nil && n.parent == nil {
			return value, true, true
		} else if n.left == nil && n.right == nil {
			if n.parent.right == n {
				n.parent.right = nil
			} else {
				n.parent.left = nil
			}
		} else if n.left == nil {
			n.value = n.right.value
			n.key = n.right.key
			n.left = n.right.left
			n.right = n.right.right

		} else if n.right == nil {
			n.value = n.left.value
			n.key = n.left.key
			n.right = n.left.right
			n.left = n.left.left

		} else if n.parent == nil {
			n.value, n.key = n.right.deleteLeft()
		} else {
			if n.parent.right == n {
				n.value, n.key = n.left.deleteLeft()
			} else {
				n.value, n.key = n.right.deleteLeft()
			}
		}
		return value, false, true
	} else if key < n.key {
		if n.left == nil {
			return nil, false, false
		}
		return n.left.delete(key)
	} else {
		if n.right == nil {
			return nil, false, false
		}
		return n.right.delete(key)
	}
}

func (n *node) deleteLeft() (interface{}, int) {
	if n.left != nil {
		return n.left.deleteLeft()
	}
	if n.right != nil {
		n.parent.right = n.right
	} else {
		n.parent.left = n.left
	}
	return n.value, n.key
}

func (n *node) preorder(ch chan *interface{}) {
	ch <- &n.value
	if n.left != nil {
		n.left.preorder(ch)
	}
	if n.right != nil {
		n.right.preorder(ch)
	}
}

func (n *node) postorder(ch chan *interface{}) {
	if n.left != nil {
		n.left.postorder(ch)
	}
	if n.right != nil {
		n.right.postorder(ch)
	}
	ch <- &n.value
}

func (n *node) inorder(ch chan *interface{}) {
	if n.left != nil {
		n.left.inorder(ch)
	}
	ch <- &n.value
	if n.right != nil {
		n.right.inorder(ch)
	}
}

// BinarySearchTree stores items in a tree as key, value pair
type BinarySearchTree struct {
	root *node
	len  int
}

// Insert a new item into the tree
func (b *BinarySearchTree) Insert(key int, value interface{}) {
	if b.root == nil {
		b.root = &node{key: key, value: value}
		b.len++
	} else if b.root.insert(key, value) {
		b.len++
	}
}

// Delete an item from the tree by the key
func (b *BinarySearchTree) Delete(key int) (interface{}, error) {
	if b.root == nil {
		return nil, errors.New("Tree is empty")
	}
	value, delRoot, isFound := b.root.delete(key)
	if isFound {
		b.len--
	} else {
		return nil, errors.New("Value cannot be found")
	}
	if delRoot {
		b.root = nil
	}
	return value, nil
}

// Iterate over the items with pre order
func (b *BinarySearchTree) IterPreorder() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		b.root.preorder(ch)
		close(ch)
	}()
	return ch
}

// Iterate over the items with post order
func (b *BinarySearchTree) IterPostorder() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		b.root.postorder(ch)
		close(ch)
	}()
	return ch
}

// Iterate over the items in order (low to high key)
func (b *BinarySearchTree) IterInorder() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		b.root.inorder(ch)
		close(ch)
	}()
	return ch
}
