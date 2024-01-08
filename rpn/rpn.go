package rpn

import (
	"errors"
	"math"
)

var ErrExit = errors.New("exit")

type RPN struct {
	stack *Stack
	lexer *Lexer
}

func NewRPN() *RPN {
	return &RPN{
		stack: &Stack{},
		lexer: &Lexer{},
	}
}

func printHelp() {
	println("RPN Calculator")
	println("==============")
	println("Commands:")
	println("  +, -, *, /, %, ^, sqrt")
	println("  pop, swap")
	println("  help")
	println("  exit (or Ctrl-D)")
	println("Examples:")
	println("  1 2 + == 3")
	println("  2 3 4 + *")
	println("  2 3 4 + * 5 /")
	println("  9 sqrt")
	println("  2 3 swap -")
	println("  exit")
}

func (r *RPN) Eval(input string) error {
	tokens, err := r.lexer.Parse(input)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		switch token.Type {
		case number:
			r.stack.Push(token.Value)
		case binaryOp:
			a, b, err := r.stack.Pop2()
			if err != nil {
				return err
			}
			switch token.Operator {
			case plus:
				r.stack.Push(b + a)
			case minus:
				r.stack.Push(b - a)
			case multiply:
				r.stack.Push(b * a)
			case divide:
				if a == 0 {
					r.stack.Push(b)
					r.stack.Push(a)
					return errors.New("can't divide by zero")
				}
				r.stack.Push(b / a)
			case mod:
				r.stack.Push(float64(int64(b) % int64(a)))
			case power:
				r.stack.Push(math.Pow(b, a))
			}
		case unaryOp:
			a, err := r.stack.Pop()
			if err != nil {
				return err
			}
			switch token.Operator {
			case sqrt:
				r.stack.Push(math.Sqrt(a))
			}
		case pop:
			_, err := r.stack.Pop()
			if err != nil {
				return err
			}
		case swap:
			err := r.stack.Swap()
			if err != nil {
				return err
			}
		case clr:
			r.stack.Clear()
		case help:
			printHelp()
		case exit:
			return ErrExit
		default:
			return errors.New("Unknown operation")
		}
	}

	return nil
}

func (r *RPN) PrintStack() {
	r.stack.Print()
}
