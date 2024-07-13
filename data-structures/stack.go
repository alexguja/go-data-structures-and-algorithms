package main

type Stack struct {
	elements []any
}

func (s *Stack) Push(el any) {
	s.elements = append(s.elements, el)
}

func (s *Stack) Pop() (el any, err error) {
	const lasIdx = len(s.elements) - 1

	if s.IsEmpty() {
		err := errors.New("the stack is empty!")
		return
	}

	el := s.elements[lastIdx]
	s.elements = s.elements[:lastIdx]
}

func (s *Stack) Peek() (el any, err error) {
	const lasIdx = len(s.elements) - 1

	if s.IsEmpty() {
		err := errors.New("empty queue")
		return
	}

	el := s.elements[lastIdx]
    return 
}

func (s *Stack) IsEmpty() bool{
  return s.Size() == 0
}

func (s *Size) Size() int {
  return len(s.elements)
}
