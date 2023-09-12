package interpreter

type Parser struct {
	Tokenizer    Tokenizer
	CurrentToken *Token
}

func (i *Parser) eat(tokenType TokenType) {
	if i.CurrentToken.Name == tokenType {
		*i.CurrentToken = i.Tokenizer.GetToken()
	} else {
		panic("Not can eat")
	}
}

func (i *Parser) factor() Expr {
	token := *i.CurrentToken
	if token.Name == INTEGER {
		i.eat(INTEGER)
		return Integer{token}
	} else {
		i.eat(PARENTHESES)
		return i.expr()
	}
}

func (i *Parser) term() Expr {
	node := i.factor()

	for i.CurrentToken.Name == MULTIPLICATION || i.CurrentToken.Name == DIVISION {
		op_token := *i.CurrentToken

		if op_token.Name == MULTIPLICATION {
			i.eat(MULTIPLICATION)
		} else {
			i.eat(DIVISION)
		}

		node = BinaryOperator{node, op_token.Name, i.factor()}
	}

	return node
}

/*
<expr> ::= <term> (("+" | "-") <term>)*
<term> ::= <factor> (("*" | "/") <factor>)*
<factor> ::= [0-9] | "(" <expr> ")"

https://bnfplayground.pauliankline.com/?bnf=<expr>%20%3A%3A%3D%20<term>%20(("%2B"%20%7C%20"-")%20<term>)*%0A<term>%20%3A%3A%3D%20<factor>%20(("*"%20%7C%20"%2F")%20<factor>)*%0A<factor>%20%3A%3A%3D%20%5B0-9%5D%20%7C%20"("%20<expr>%20")"&name=
*/
// https://bnfplayground.pauliankline.com/?bnf=<expr>%20%3A%3A%3D%20<term>%20(("%2B"%20%7C%20"-")%20<term>)*%0A<term>%20%3A%3A%3D%20<factor>%20(("*"%20%7C%20"%2F")%20<factor>)*%0A<factor>%20%3A%3A%3D%20%5B0-9%5D%20%7C%20"("%20<expr>%20")"&name=
func (i *Parser) expr() Expr {
	if i.CurrentToken.Value == nil {
		*i.CurrentToken = i.Tokenizer.GetToken()
	}
	node := i.term()

	for i.CurrentToken.Name == PLUS || i.CurrentToken.Name == MINUS {
		op_token := *i.CurrentToken

		if op_token.Name == PLUS {
			i.eat(PLUS)
		} else {
			i.eat(MINUS)
		}
		node = BinaryOperator{node, op_token.Name, i.term()}
	}

	return node
}

func (i *Parser) Parse() Expr {
	return i.expr()
}
