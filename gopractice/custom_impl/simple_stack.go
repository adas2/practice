package custom

import "fmt"

// simple stack implementation

// Sstack simple stack impl
type Sstack []int

// InitStack : input max capacity (optional)
func InitStack(cap int) *Sstack {
	stck := make(Sstack, 0, cap)
	return &stck
}

// Push to the top of stack
func (s *Sstack) Push(val int) {
	// todo: error case when stack is full
	// end of the slice is top of stack
	*s = append(*s, val)
}

// Pop from top, dels the element
func (s *Sstack) Pop() (int, error) {
	// empty case
	if len(*s) == 0 {
		return -1, fmt.Errorf("stack empty")
	}
	// adjust len
	n := len(*s)
	res := (*s)[n-1]
	*s = (*s)[:n-1]
	return res, nil
}

// IsEmpty tells if stack is empty
func (s *Sstack) IsEmpty() bool {
	if len(*s) == 0 {
		return false
	}
	return true
}

// Peek returns top element without popping it
func (s *Sstack) Peek() int {
	n := len(*s)
	if n == 0 {
		return -1 //error
	}
	return (*s)[n-1]
}
