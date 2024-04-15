package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	items []interface{}
	len   uint32
	lock  sync.Mutex
}

func New() *Queue {
	stack := new(Queue)
	stack.items = make([]interface{}, 0)
	stack.len = 0

	return stack
}

func (s *Queue) Length() uint32 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len
}

func (s *Queue) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len == 0
}

func (s *Queue) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.len != 0 {
		return s.items[0]
	}

	return nil
}

func (s *Queue) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, val)
	s.len++
}

func (s *Queue) Pop() interface{} {
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
