package stack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()

	// Validate Length
	if !s.IsEmpty() || s.len != 0 || s.Length() != 0 {
		t.Error()
	}

	s.Push("First Value")
	s.Push(2)
	s.Push(3.0)

	// Validate Inserted Value
	if s.items[0] != 3.0 || s.items[1] != 2 || s.items[2] != "First Value" {
		fmt.Println(s.items...)
		t.Error()
	}

	if s.Length() != 3 {
		t.Error()
	}

	// Remove First Index Value
	removedValue := s.Pop()
	if removedValue != 3.0 {
		fmt.Println(removedValue)
		t.Error()
	}

	if s.items[0] != 2 || s.items[1] != "First Value" || s.Length() != 2 {
		t.Error()
	}

	// Check First Value
	firstValue := s.Peek()
	if firstValue != 2 {
		t.Error()
	}
}
