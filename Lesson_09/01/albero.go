package main

import (
	"fmt"
	"math/rand"
) 

type bitreeNode struct {
	left *bitreeNode
	right *bitreeNode
	val int
}

type bitree struct {
	root * bitreeNode
}

func main () {
	var slice = []int{}
	var MAX_ELEMENTS = 12
	for i := 0; i < MAX_ELEMENTS; i++ {
		slice = append(slice, rand.Intn(100))
	}
	fmt.Println("Slice: ", slice)
	var tree = arr2tree(slice, 0)
	fmt.Println("Tree: ")
	printTree(tree, 0)
}

func arr2tree (a []int, i int) ( root * bitreeNode ) {
	if i < len(a) {
		root = &bitreeNode{val: a[i]}
		root.left = arr2tree(a, 2*i+1)
		root.right = arr2tree(a, 2*i+2)
	}
	return
}

func printTree (root * bitreeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print("  ")
	}
	if root != nil {
		fmt.Print("*")
		fmt.Println(root.val)
		if(root.left != nil || root.right != nil) {
			printTree(root.left, spaces+1)
			printTree(root.right, spaces+1)
		}
	} else {
		fmt.Print("*")
		fmt.Println()
	}
}
