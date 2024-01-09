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

func (l *Lexer) parseNumber(negative bool) (*Token, error) {
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
		return nil, err
	}
	if negative {
		numStr = "-" + numStr
		num = -num
	}

	return &Token{number, 0, num, numStr}, nil
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
		return &Token{unaryOp, sqrt, 0, word}, nil
	case "dec":
		return &Token{dec, 0, 0, word}, nil
	case "bin":
		return &Token{bin, 0, 0, word}, nil
	case "hex":
		return &Token{hex, 0, 0, word}, nil
	case "pi":
		return &Token{number, 0, math.Pi, word}, nil
	case "pop":
		return &Token{pop, 0, 0, word}, nil
	case "swap":
		return &Token{swap, 0, 0, word}, nil
	case "clr":
		return &Token{clr, 0, 0, word}, nil
	case "help":
		return &Token{help, 0, 0, word}, nil
	case "exit":
		return &Token{exit, 0, 0, word}, nil
	default:
		return nil, fmt.Errorf("Unknown input: %s", word)
	}
}

func (l *Lexer) Parse(input string) ([]*Token, error) {
	var err error
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
			token = &Token{binaryOp, plus, 0, string(char)}
		case '-':
			char = l.input.Eat()
			if unicode.IsNumber(char) {
				token, err = l.parseNumber(true)
				if err != nil {
					return nil, err
				}
			} else {
				token = &Token{binaryOp, minus, 0, string(char)}
			}
		case '*':
			l.input.Eat()
			token = &Token{binaryOp, multiply, 0, string(char)}
		case '/':
			l.input.Eat()
			token = &Token{binaryOp, divide, 0, string(char)}
		case '%':
			l.input.Eat()
			token = &Token{binaryOp, mod, 0, string(char)}
		case '^':
			l.input.Eat()
			token = &Token{binaryOp, power, 0, string(char)}
		case ' ':
			l.input.Eat()
			continue
		case 0x0a:
			l.input.Eat()
			continue
		case 0x04: // Ctrl-D
			token = &Token{exit, 0, 0, "^D"}
		case '?':
			l.input.Eat()
			token = &Token{help, 0, 0, string(char)}
		default:
			if unicode.IsNumber(char) {
				token, err = l.parseNumber(false)
				if err != nil {
					return nil, err
				}
			} else if unicode.IsLetter(char) {
				token, err = l.parseWord()
				if err != nil {
					return nil, err
				}
			} else {
				return nil, fmt.Errorf("Invalid input: %c", char)
			}
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
