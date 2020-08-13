package maze

import (
	"sync"
)

//Stack is a basic stack implementation for cells which supports concurrency
type Stack struct {
	lock  sync.Mutex
	slice []*Cell
}

//Push adds an element to the top of the stack
func (s *Stack) Push(c *Cell) {
	s.lock.Lock()
	s.slice = append(s.slice, c)
	s.lock.Unlock()
}

//Pop removes an element from the top of the stack and returns it
func (s *Stack) Pop() *Cell {

	s.lock.Lock()
	l := len(s.slice) - 1
	pop := s.slice[l]
	s.slice = append(s.slice[:l], s.slice[l+1:]...)
	s.lock.Unlock()
	return pop
}

//IsEmpty returns whether the stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.slice) == 0
}
