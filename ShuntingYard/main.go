package shuntingyard

import (
	"strconv"
	"strings"

	"github.com/victoralmeida428/EquacaoPolonesa/queue"
	"github.com/victoralmeida428/EquacaoPolonesa/stack"
)

// Precedence and associativity of operators
var Precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
var leftAssociative = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

func IsOperator(token string) bool {
	_, exists := Precedence[token]
	return exists
}

func isFunction(token string) bool {
	// Define supported functions here
	functions := map[string]bool{"sin": true, "cos": true, "tan": true, "log": true}
	return functions[token]
}

// ShuntingYard implements the algorithm
func ShuntingYard(expression string) string {
	tokens := strings.Fields(expression) // Split expression into tokens
	outputQueue := queue.New()           // Output queue
	operatorStack := stack.Stack{}       // Operator stack

	for _, token := range tokens {
		if isNumber(token) {
			// If the token is a number, add it to the output queue
			outputQueue.Add(token)
		} else if isFunction(token) {
			// If the token is a function, push it onto the operator stack
			operatorStack.Push(token)
		} else if IsOperator(token) {
			// If the token is an operator
			for !operatorStack.IsEmpty() {
				top := operatorStack.Peek()
				if IsOperator(top) &&
					(Precedence[top] > Precedence[token] ||
						(Precedence[top] == Precedence[token] && leftAssociative[token])) {
					outputQueue.Add(operatorStack.Pop())
				} else {
					break
				}
			}
			operatorStack.Push(token)
		} else if token == "(" {
			// Push left parenthesis onto the stack
			operatorStack.Push(token)
		} else if token == ")" {
			// Pop from the stack to the output queue until left parenthesis is found
			for !operatorStack.IsEmpty() && operatorStack.Peek() != "(" {
				outputQueue.Add(operatorStack.Pop())
			}
			if operatorStack.IsEmpty() {
				panic("Mismatched parentheses")
			}
			operatorStack.Pop() // Discard the left parenthesis
		}
	}

	// Pop remaining operators into the output queue
	for !operatorStack.IsEmpty() {
		top := operatorStack.Pop()
		if top == "(" || top == ")" {
			panic("Mismatched parentheses")
		}
		outputQueue.Add(top)
	}

	// Convert queue to space-separated string
	return strings.Join(toStringSlice(outputQueue.GetData()), " ")
}

// Helper functions
func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func toStringSlice(data []interface{}) []string {
	result := make([]string, len(data))
	for i, v := range data {
		result[i] = v.(string)
	}
	return result
}
