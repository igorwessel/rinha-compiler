package interpreter

type Expr interface{}

type Integer struct {
	token Token
}

func (i *Integer) Value() int {
	return i.token.Value.(int)
}

type BinaryOperator struct {
	left  Expr
	op    TokenType
	right Expr
}
