package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New()

	// Testa enfileiramento
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	expectedData := []interface{}{1, 2, 3}
	if got := queue.GetData(); !equal(got, expectedData) {
		t.Errorf("GetData() = %v; want %v", got, expectedData)
	}

	// Testa o elemento frontal
	if front := queue.Front(); front != 1 {
		t.Errorf("Front() = %v; want %v", front, 1)
	}

	// Testa desenfileiramento
	queue.Remove()
	expectedData = []interface{}{2, 3}
	if got := queue.GetData(); !equal(got, expectedData) {
		t.Errorf("After Remove(), GetData() = %v; want %v", got, expectedData)
	}

	// Testa novamente o elemento frontal
	if front := queue.Front(); front != 2 {
		t.Errorf("After Remove(), Front() = %v; want %v", front, 2)
	}

	// Testa esvaziar a fila
	queue.Remove()
	queue.Remove()
	if front := queue.Front(); front != nil {
		t.Errorf("After emptying, Front() = %v; want nil", front)
	}

	if got := queue.GetData(); len(got) != 0 {
		t.Errorf("After emptying, GetData() = %v; want empty queue", got)
	}
}

// Função auxiliar para comparar fatias
func equal(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
