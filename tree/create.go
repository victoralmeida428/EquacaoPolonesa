package tree

import shuntingyard "equacao/ShuntingYard"

func (tree *Tree) buildTree(){
	tree.Root = tree.createNode()
}

func (tree *Tree) createNode() *Node{
	if tree.data == nil {
		return nil
	}

	token := tree.data.Pop()

	node := &Node{
		Data: token,
	}
	if shuntingyard.IsOperator(token){
		node.Left = tree.createNode()
		node.Rigth = tree.createNode()
	}
	
	return node	
}