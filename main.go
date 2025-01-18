package main

import (
	"fmt"

	"numerical_expression/tree"
)

func main() {
	exp := "( ( ( 2 ^ 3 ) * ( 5 + 2 ) ) / ( 7 - ( 2 * 3 ) ) ) + ( ( 8 + 2 ) * ( 4 ^ 2 ) ) - ( ( 10 / 2 ) + 3 )"
	t := tree.New(exp)
	fmt.Println("Expressão: ", exp)
	fmt.Println("Árvore criada com sucesso!")
	t.Print()
	fmt.Println("Resultado: ", t.Calculate())


	exp = "10 % 3"
	t = tree.New(exp)
	fmt.Println("Expressão: ", exp)
	fmt.Println("Árvore criada com sucesso!")
	t.Print()
	fmt.Println("Resultado: ", t.Calculate())
}
