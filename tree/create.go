package tree

import shuntingyard "github.com/victoralmeida428/EquacaoPolonesa/ShuntingYard"

func (tree *Tree) buildTree() {
	tree.Root = tree.createNode()
}

func (tree *Tree) createNode() *Node {
	if tree.data == nil {
		return nil
	}

	token := tree.data.Pop()

	node := &Node{
		Data: token,
	}
	if shuntingyard.IsOperator(token) {
		node.Left = tree.createNode()
		node.Right = tree.createNode()
	}

	return node
}
