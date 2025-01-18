package stack

import "testing"

func TestPush(t *testing.T) {
	stack := New()

	stack.Push("A")
	stack.Push("B")

	if len(*stack) != 2 {
		t.Errorf("Expected stack size 2, but got %d", len(*stack))
	}

	if (*stack)[0] != "A" || (*stack)[1] != "B" {
		t.Errorf("Push did not add elements correctly. Stack contents: %v", *stack)
	}
}

func TestPop(t *testing.T) {
	stack := New()

	stack.Push("A")
	stack.Push("B")
	element := stack.Pop()

	if element != "B" {
		t.Errorf("Expected popped element to be 'B', but got '%s'", element)
	}

	if len(*stack) != 1 {
		t.Errorf("Expected stack size 1 after pop, but got %d", len(*stack))
	}

	element = stack.Pop()
	if element != "A" {
		t.Errorf("Expected popped element to be 'A', but got '%s'", element)
	}

	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all elements")
	}
}

func TestPeek(t *testing.T) {
	stack := New()

	if stack.Peek() != "" {
		t.Errorf("Expected empty stack Peek to return '', but got '%s'", stack.Peek())
	}

	stack.Push("A")
	stack.Push("B")

	if stack.Peek() != "B" {
		t.Errorf("Expected top element to be 'B', but got '%s'", stack.Peek())
	}

	stack.Pop()
	if stack.Peek() != "A" {
		t.Errorf("Expected top element to be 'A', but got '%s'", stack.Peek())
	}
}

func TestIsEmpty(t *testing.T) {
	stack := New()

	if !stack.IsEmpty() {
		t.Errorf("Expected new stack to be empty")
	}

	stack.Push("A")
	if stack.IsEmpty() {
		t.Errorf("Expected non-empty stack after push")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("Expected empty stack after popping all elements")
	}
}
