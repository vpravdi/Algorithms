//find closest value in BST
package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {

	return tree.Value
}
