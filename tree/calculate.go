package tree

import (
	"fmt"
	"strconv"
)

func (tree *Tree) evaluate(node *Node) float64 {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		value, err := strconv.ParseFloat(node.Data, 64)
		if err != nil {
			panic(fmt.Sprintf("Error to convert '%s' for a number: %v", node.Data, err))
		}
		return value
	}

	leftValue := tree.evaluate(node.Left)
	rightValue := tree.evaluate(node.Right)
	
	switch node.Data {
	case "+":
		return rightValue+leftValue
	case "-":
		return rightValue - leftValue
	case "*":
		return leftValue * rightValue
	case "/":
		if leftValue == 0{
			panic("")
		}
		return rightValue/leftValue
	default:
		panic("Operation not found")
	}
}

func (tree *Tree) Calculate()float64{
	return tree.evaluate(tree.Root)
}
