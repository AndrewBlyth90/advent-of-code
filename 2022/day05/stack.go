package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
}

func (s Stack) Empty() bool {
	return s.stack.Len() == 0
}

func (s Stack) Peek() (interface{}, error) {
	if s.stack.Len() > 0 {
		return s.stack.Front().Value, nil
	} else {
		return nil, fmt.Errorf("Peek Error: Stack is empty")
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.stack.Len() > 0 {
		element := s.stack.Front()
		return s.stack.Remove(element), nil
	} else {
		return nil, fmt.Errorf("Pop Error: Stack is empty")
	}
}

func (s *Stack) Push(element interface{}) interface{} {
	s.stack.PushFront(element)
	return element
}

func NewStack() Stack {
	return Stack{
		stack: list.New(),
	}
}