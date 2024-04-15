package stack

import (
	"fmt"
	"sync"
)

type Stack struct {
	items []interface{}
	len   uint32
	lock  sync.Mutex
}

func New() *Stack {
	stack := new(Stack)
	stack.items = make([]interface{}, 0)
	stack.len = 0

	return stack
}

func (s *Stack) Length() uint32 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len == 0
}

func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.len != 0 {
		return s.items[0]
	}

	return nil
}

func (s *Stack) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	firstValue := make([]interface{}, 0)
	firstValue = append(firstValue, val)

	s.items = append(firstValue, s.items...)
	s.len++
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.len == 0 {
		return ""
	}

	removedElement, updatedValue := s.items[0], s.items[1:]
	s.items = updatedValue
	s.len--

	fmt.Println("Removed Element:", removedElement)
	fmt.Println("Updated Value:", updatedValue)

	return removedElement
}
