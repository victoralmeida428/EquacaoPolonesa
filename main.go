package main

import (
	"fmt"

	"github.com/victoralmeida428/EquacaoPolonesa/tree"
)

func main() {
	exp := "( ( 2 + 3 ) * ( 2 + 3 ) ) / 5 + 1 "
	t := tree.New(exp)
	fmt.Println("Expressão: ", exp)
	fmt.Println("Árvore criada com sucesso!")
	t.Print()
	fmt.Println("Resultado: ", t.Calculate())
}
