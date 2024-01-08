package rpn

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

type Lexer struct {
	input *Input
}

func (l *Lexer) parseNumber() (float64, error) {
	numStr := ""
	char := l.input.NextChar()
	for {
		if !unicode.IsNumber(char) && char != '.' {
			break
		}

		numStr += string(char)
		char = l.input.Eat()
	}

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func (l *Lexer) parseWord() (*Token, error) {
	word := ""
	char := l.input.NextChar()
	for {
		if !unicode.IsLetter(char) {
			break
		}

		word += string(char)
		char = l.input.Eat()
	}

	switch word {
	case "sqrt":
		return &Token{unaryOp, sqrt, 0}, nil
	case "dec":
		return &Token{dec, 0, 0}, nil
	case "bin":
		return &Token{bin, 0, 0}, nil
	case "hex":
		return &Token{hex, 0, 0}, nil
	case "pi":
		return &Token{number, 0, math.Pi}, nil
	case "pop":
		return &Token{pop, 0, 0}, nil
	case "swap":
		return &Token{swap, 0, 0}, nil
	case "clr":
		return &Token{clr, 0, 0}, nil
	case "help":
		return &Token{help, 0, 0}, nil
	case "exit":
		return &Token{exit, 0, 0}, nil
	default:
		return nil, fmt.Errorf("Unknown input: %s", word)
	}
}

func (l *Lexer) Parse(input string) ([]*Token, error) {
	var err error
	var num float64
	var token *Token
	l.input = NewInput(input)
	tokens := []*Token{}
	for {
		char := l.input.NextChar()
		if char == 0 {
			break
		}

		switch char {
		case '+':
			l.input.Eat()
			token = &Token{binaryOp, plus, 0}
		case '-':
			char = l.input.Eat()
			if unicode.IsNumber(char) {
				num, err = l.parseNumber()
				token = &Token{number, 0, -num}
				if err != nil {
					return nil, err
				}
			} else {
				token = &Token{binaryOp, minus, 0}
			}
		case '*':
			l.input.Eat()
			token = &Token{binaryOp, multiply, 0}
		case '/':
			l.input.Eat()
			token = &Token{binaryOp, divide, 0}
		case '%':
			l.input.Eat()
			token = &Token{binaryOp, mod, 0}
		case '^':
			l.input.Eat()
			token = &Token{binaryOp, power, 0}
		case ' ':
			l.input.Eat()
			continue
		case 0x0a:
			l.input.Eat()
			continue
		case 0x04: // Ctrl-D
			token = &Token{exit, 0, 0}
		case '?':
			l.input.Eat()
			token = &Token{help, 0, 0}
		default:
			if unicode.IsNumber(char) {
				num, err = l.parseNumber()
				if err != nil {
					return nil, err
				}
				token = &Token{number, 0, num}
			} else if unicode.IsLetter(char) {
				token, err = l.parseWord()
				if err != nil {
					return nil, err
				}
			} else {
				return nil, fmt.Errorf("Unknown input: %c", char)
			}
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
