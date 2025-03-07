package rpn

import (
	"fmt"
	"math"
)

var ErrNotEnoughValuesTmpl = "Need at least %d values on the stack to perform this operation"

type Stack struct {
	values []float64
}

func (s *Stack) Push(value float64) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (float64, error) {
	count := s.Len()
	if count == 0 {
		return 0, fmt.Errorf(ErrNotEnoughValuesTmpl, 1)
	}
	value := s.values[count-1]
	s.values = s.values[:count-1]
	return value, nil
}

func (s *Stack) Pop2() (float64, float64, error) {
	count := s.Len()
	if count < 2 {
		return 0, 0, fmt.Errorf(ErrNotEnoughValuesTmpl, 2)
	}
	a := s.values[count-1]
	b := s.values[count-2]
	s.values = s.values[:count-2]
	return a, b, nil
}

func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) Swap() error {
	count := s.Len()
	if count < 2 {
		return fmt.Errorf(ErrNotEnoughValuesTmpl, 2)
	}
	a := s.values[count-1]
	b := s.values[count-2]
	s.values[count-1] = b
	s.values[count-2] = a
	return nil
}

func (s *Stack) Clear() {
	s.values = []float64{}
}

func (s *Stack) Print() {
	for _, value := range s.values {
		if value == math.Trunc(value) {
			fmt.Printf("%d\n", int64(value))
		} else {
			fmt.Printf("%f\n", value)
		}
	}
}
