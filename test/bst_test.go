package main

import ("fmt";"ds/bst")

func main() {
	bst := bst.NewBinaryTree()
	bst.Add(4)
	bst.Add(5)
	bst.Add(1)
	bst.Add(3)
	bst.Add(2)
	bst.Add(12)
	bst.Add(9)
	
	//this way you don't have to recalculate traversals unless tree
	//changes if you need them again.	
	trav := bst.Traverse("in")
	fmt.Println("in order traversal")
	for e := trav.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
	
	trav = bst.Traverse("pre")
	fmt.Println("pre order traversal")
	for e := trav.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	trav = bst.Traverse("post")
	fmt.Println("post order traversal")
	for e := trav.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
