package shuntingyard

import "testing"

func TestShuntingYard(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
	}{
		// Testes simples
		{"2 + 3", "2 3 +"},
		{"2 * 3", "2 3 *"},
		{"2 + 3 * 4", "2 3 4 * +"},
		{"( 2 + 3 ) * 4", "2 3 + 4 *"},
		{"2 * ( 3 + 4 )", "2 3 4 + *"},

		// Testes com múltiplos operadores
		{"2 + 3 - 4", "2 3 + 4 -"},
		{"2 * 3 / 4", "2 3 * 4 /"},
		{"2 + 3 * 4 - 5", "2 3 4 * + 5 -"},
		{"( 2 + 3 ) * ( 4 - 5 )", "2 3 + 4 5 - *"},

		// Testes com parênteses aninhados
		{"( 2 + ( 3 * 4 ) ) - 5", "2 3 4 * + 5 -"},
		{"( ( 2 + 3 ) * 4 ) / 5", "2 3 + 4 * 5 /"},

		// Testes com funções (exemplo com `sin`)
		{"sin ( 90 )", "90 sin"},
		{"2 + sin ( 90 )", "2 90 sin +"},
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result := ShuntingYard(test.expression)
			if result != test.expected {
				t.Errorf("ShuntingYard(%q) = %q; expected %q", test.expression, result, test.expected)
			}
		})
	}
}

func TestShuntingYardWithErrors(t *testing.T) {
	tests := []struct {
		expression string
		shouldPanic bool
	}{
		// Testes de erros com parênteses
		{"( 2 + 3", true},   // Falta parêntese direito
		{"2 + 3 )", true},   // Falta parêntese esquerdo
		{"2 + ( 3 * 4 ))", true}, // Parênteses extras
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !test.shouldPanic {
						t.Errorf("ShuntingYard(%q) inesperadamente gerou pânico: %v", test.expression, r)
					}
				} else {
					if test.shouldPanic {
						t.Errorf("ShuntingYard(%q) deveria ter gerado pânico, mas não gerou", test.expression)
					}
				}
			}()

			_ = ShuntingYard(test.expression)
		})
	}
}
