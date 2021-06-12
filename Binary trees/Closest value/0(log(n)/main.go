//find closest value in BST
//0(log(n)) time space 0(log(n))
package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {

	return tree.Value
}
