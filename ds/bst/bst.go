package bst

import ("fmt")

/*
 * Class to construc a binary search tree. 
 * Does not have a remove method yet but it will
 * be added soon. Insertions and lookups should
 * be O(log n) if the tree is mostly balanced.
 */

// Nonexported struct to build tree
type node struct {
    data int
    left *node
    right *node
}

// Contains data about the binary tree
type BinarySearchTree struct {
    head *node
    size int
}

// Initializes a new binary tree
func NewBinaryTree() *BinarySearchTree {
    tree := new(BinarySearchTree)
    tree.head = new(node)
    return tree
}

// Nonexported method to iterate over tree
func add(head *node, item int) {
   if head.data > item {
       if head.left == nil {
           n := new(node)
           n.data = item
           head.left = n
       } else {
           add(head.left, item)
       }
   } else {
       if head.right == nil {
           n := new(node)
           n.data = item
           head.right = n
        } else {
           add(head.right, item)
        }
   }
}

// Adds an integer into the tree
func (tree *BinarySearchTree) Add(item int) {
    if tree.size == 0 {
        n := new(node)
        n.data = item
        tree.head = n
        tree.size++
    } else {
        add(tree.head, item)
        tree.size++
    }
}

// Returns the size of the tree
func (tree BinarySearchTree) Size() int {
    return tree.size;
}

// Nonexported method to help add an item
func contains(head *node, item int) bool{
    if head.data == item {
        return true
    } else if head.data > item {
        if head.left == nil {
            return false
        } else {
            contains(head.left, item)
        }
    } else {
        if head.right == nil {
            return false
        } else {
            contains(head.right, item)
        }
    }
    return false
}

// Checks whether or not the tree cotains the given item
func (tree BinarySearchTree) Contains(item int) bool {
    return contains(tree.head, item)
}

// Helper method to print out the tree
func printHelper(n *node) {
    if n == nil {
        return
    }
    fmt.Printf("%v\n", n.data)
    printHelper(n.left)
    printHelper(n.right)
}

// Will print out the tree in pre-order 
func (tree BinarySearchTree) Print() {
   printHelper(tree.head) 
}
