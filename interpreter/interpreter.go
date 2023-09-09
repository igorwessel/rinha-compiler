package interpreter

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
	if token.Name == INTEGER {
		i.eat(INTEGER)
		return token.Value.(int)
	} else {
		i.eat(PARENTHESES)
		return i.Expr()
	}
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

/*
<expr> ::= <term> (("+" | "-") <term>)*
<term> ::= <factor> (("*" | "/") <factor>)*
<factor> ::= [0-9] | "(" <expr> ")"

https://bnfplayground.pauliankline.com/?bnf=<expr>%20%3A%3A%3D%20<term>%20(("%2B"%20%7C%20"-")%20<term>)*%0A<term>%20%3A%3A%3D%20<factor>%20(("*"%20%7C%20"%2F")%20<factor>)*%0A<factor>%20%3A%3A%3D%20%5B0-9%5D%20%7C%20"("%20<expr>%20")"&name=
*/
// https://bnfplayground.pauliankline.com/?bnf=<expr>%20%3A%3A%3D%20<term>%20(("%2B"%20%7C%20"-")%20<term>)*%0A<term>%20%3A%3A%3D%20<factor>%20(("*"%20%7C%20"%2F")%20<factor>)*%0A<factor>%20%3A%3A%3D%20%5B0-9%5D%20%7C%20"("%20<expr>%20")"&name=
func (i *Interpreter) Expr() int {
	if i.CurrentToken.Value == nil {
		*i.CurrentToken = i.Tokenizer.GetToken()
	}
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
