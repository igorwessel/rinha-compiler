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

