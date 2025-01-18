package tree

import (
	shuntingyard "equacao/ShuntingYard"
	"equacao/stack"
	"fmt"
	"strings"
)

type ITree interface {
	GetRoot() *Node
	Print()
	Calculate() float64
}

type Tree struct {
	Exp  string `json:"expression"`
	data *stack.Stack
	Root *Node
}

func (tree *Tree) convertExp() {
	exp := shuntingyard.ShuntingYard(tree.Exp)
	tokens := strings.Fields(exp)
	for _, token := range tokens {
		tree.data.Push(token)
	}
}

func (tree *Tree) GetRoot() *Node {
	return tree.Root
}
func (tree *Tree) printTree(node *Node, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	// Formata a linha para o nó atual
	arrow := "├──"
	if !isLeft {
		arrow = "└──"
	}

	fmt.Println(prefix + arrow + node.Data)

	// Atualiza o prefixo para os filhos
	childPrefix := prefix
	if isLeft {
		childPrefix += "│   " // Linha vertical para a continuidade da subárvore
	} else {
		childPrefix += "    " // Espaço vazio, pois é o último nó
	}

	// Chama recursivamente para os filhos
	tree.printTree(node.Left, childPrefix, true)
	tree.printTree(node.Right, childPrefix, false)
}

func (tree *Tree) Print() {
	tree.printTree(tree.Root, "   ", false)
}

func New(exp string) ITree {
	tree := Tree{Exp: exp, data: stack.New()}
	tree.convertExp()
	tree.buildTree()
	return &tree
}
