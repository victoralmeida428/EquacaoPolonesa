package tree

import (
	"testing"
)

func TestTreeBuildAndEvaluate(t *testing.T) {
	tests := []struct {
		expression  string   // Notação pós-fixa (RPN)
		expected    float64  // Resultado esperado
	}{
		// Testes simples
		{"2 + 3", 5},
		{"2 * 3", 6},
		{"2 * ( 3 + 4 )", 14},
		{"( 2 + 3 ) * 4", 20},
		{"2 + 3 * 4", 14},

		// Testes com múltiplos operadores
		{"2 + 3 - 4", 1},
		{"2 * 3 / 4", 1.5},
		{"2 + 3 * 4 - 5", 9},

		// Testes com parênteses aninhados
		{"( 2 + 3 ) * ( 4 - 5 )", -5},
		{"( ( 2 + 3 ) * 4 ) / 5", 4},
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			// Cria uma pilha com os tokens da expressão
			// Constrói a árvore
			tree := New(test.expression)
			// Avalia a árvore
			result := tree.Calculate()
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
