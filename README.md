
# Expression Calculator
This project implements a calculator for mathematical expressions using binary trees to represent and evaluate infix expressions. The input supports standard mathematical expressions, such as "( ( 2 + 3 ) * ( 2 + 3 ) ) / 5 + 1".

## Features
- Converts infix expressions to Reverse Polish Notation (RPN) using the Shunting Yard algorithm.
- Constructs a binary tree to represent the expression.
- Evaluates the expression by traversing the binary tree.

## Example Usage
```go
package main

import (
	"fmt"

	"github.com/victoralmeida428/EquacaoPolonesa/tree"
)

func main() {
	exp := "( ( 2 + 3 ) * ( 2 + 3 ) ) / 5 + 1 "
	t := tree.New(exp)
	fmt.Println("Expression: ", exp)
	fmt.Println("Tree successfully created!")
	t.Print()
	fmt.Println("Result: ", t.Calculate())
}

```
#### Expected Output
For the expression:
```scss
( ( 2 + 3 ) * ( 2 + 3 ) ) / 5 + 1
```
The program will output:
```scss
Expression:  ( ( 2 + 3 ) * ( 2 + 3 ) ) / 5 + 1 
Tree successfully created!
   └──+
       ├──1
       └──/
           ├──5
           └──*
               ├──+
               │   ├──3
               │   └──2
               └──+
                   ├──3
                   └──2
Result:  5

```
