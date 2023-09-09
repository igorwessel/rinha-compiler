package interpreter

import (
	"fmt"
)

type Interpreter struct {
	Tokenizer    Tokenizer
	CurrentToken *Token
}

func (i *Interpreter) eat(tokenType TokenType) {
	if i.CurrentToken.Name == tokenType {
		*i.CurrentToken = i.Tokenizer.GetToken()
	} else {
		panic("Not can eat")
	}
}

func (i *Interpreter) factor() int {
	token := *i.CurrentToken
	fmt.Println(token)
	i.eat(INTEGER)
	return token.Value.(int)
}

func (i *Interpreter) term() int {

	result := i.factor()
	for i.CurrentToken.Name == MULTIPLICATION || i.CurrentToken.Name == DIVISION {
		op_token := *i.CurrentToken

		if op_token.Name == MULTIPLICATION {
			i.eat(MULTIPLICATION)
			result *= i.factor()
		} else {
			i.eat(DIVISION)
			result /= i.factor()
		}
	}

	return result
}

func (i *Interpreter) parent() int {
	result := i.term()

	for i.CurrentToken.Name == PLUS || i.CurrentToken.Name == MINUS {
		op_token := *i.CurrentToken

		if op_token.Name == PLUS {
			i.eat(PLUS)
			result += i.term()
		} else {
			i.eat(MINUS)
			result -= i.term()
		}

	}

	return result
}

func (i *Interpreter) Expr() int {
	*i.CurrentToken = i.Tokenizer.GetToken()
	result := i.parent()

	for i.CurrentToken.Name == PARENTHESES {
		fmt.Println(i.CurrentToken)
		op_token := *i.CurrentToken

		if op_token.Name == PARENTHESES {
			i.eat(PARENTHESES)
			result += i.parent()
		}
	}

	return result
}
