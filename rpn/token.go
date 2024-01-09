package rpn

type TokenType int

const (
	number TokenType = iota
	binaryOp
	unaryOp
	dec
	bin
	hex
	pop
	swap
	clr
	help
	exit
)

type Operator int

const (
	plus Operator = iota
	minus
	multiply
	divide
	mod
	power
	sqrt
)

type Token struct {
	Type     TokenType
	Operator Operator
	Value    float64
	Original string
}
