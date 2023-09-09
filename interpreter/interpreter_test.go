package interpreter

import (
	"testing"
)

func TestBasicPlus(t *testing.T) {

	text := "1+1"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 2 {
		t.Fatalf(`should be 2 in basic sum expression like 1+1, we receive %q`, result)
	}
}

func TestBasicMinus(t *testing.T) {

	text := "1-1"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 0 {
		t.Fatalf(`should be 0 in basic minus expression like 1-1, we receive %q`, result)
	}
}

func TestBasicMul(t *testing.T) {

	text := "2*2"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 4 {
		t.Fatalf(`should be 4 in mult expression like 2*2, we receive %q`, result)
	}
}
func TestBasicDiv(t *testing.T) {

	text := "2/2"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 1 {
		t.Fatalf(`should be 1 in div expression like 2/2, we receive %q`, result)
	}
}

func TestIgnoreWhitespace(t *testing.T) {

	text := "2          *             2"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 4 {
		t.Fatalf(`should ignore whitespace in expression like 2    *         2, we receive %q`, result)
	}
}

func TestPrecendenceOperatorMul(t *testing.T) {

	text := "2+2*2"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 6 {
		t.Fatalf(`should be 6 in expression like 2+2*2, we receive %q`, result)
	}
}

func TestPrecendenceOperatorDiv(t *testing.T) {

	text := "2+2/2"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 3 {
		t.Fatalf(`should be 3 in div expression like 2+2/2, we receive %q`, result)
	}

}

func TestPrecendeOperatorMulDiv(t *testing.T) {

	text := "2+1-2*2/2"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 1 {
		t.Fatalf(`should be 1 in expression like 2+1-2*2/2, we receive %q`, result)
	}
}

func TestPrecendenceParentheses(t *testing.T) {

	text := "2+(5-1)"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 6 {
		t.Fatalf(`should be 6 in precedence parentheses expression like 2+(5-1), we receive %q`, result)
	}
}

func TestComplexExpression(t *testing.T) {

	text := "2+(5-(1+3))"
	interpreter := Interpreter{Tokenizer{text, 0, nil}, &Token{}}
	result := interpreter.Expr()

	if result != 3 {
		t.Fatalf(`should be 3 in complex expression like 2+(5-(1+3)), we receive %q`, result)
	}

}
