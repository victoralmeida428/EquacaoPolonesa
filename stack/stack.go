package stack

type Stack []string

// Push an element onto the stack
func (s *Stack) Push(element string) {
	*s = append(*s, element)
}

// Pop an element from top of the stack
func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element
}

// Peek the top element of the stack
func (s *Stack) Peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}


func New()*Stack{
	return &Stack{}
}