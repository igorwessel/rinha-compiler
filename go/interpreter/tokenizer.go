package interpreter

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type TokenType string

const (
	INTEGER        TokenType = "INTEGER"
	PLUS           TokenType = "PLUS"
	MINUS          TokenType = "MINUS"
	EOF            TokenType = "EOF"
	MULTIPLICATION TokenType = "MULTIPLICATION"
	PARENTHESES    TokenType = "PARENTHESES"
	DIVISION       TokenType = "DIVISION"
)

type Token struct {
	Name  TokenType
	Value any
}

type Tokenizer struct {
	Text        string
	Pos         int
	CurrentChar *byte
}

func (t *Tokenizer) Next() {
	t.Pos += 1
	if t.Pos > len(t.Text)-1 {
		t.CurrentChar = nil
	} else {
		*t.CurrentChar = t.Text[t.Pos]
	}
}

func (t *Tokenizer) digit() int {
	result := ""

	for t.CurrentChar != nil && regexp.MustCompile(`\d`).Match([]byte{*t.CurrentChar}) {
		result += string(*t.CurrentChar)
		t.Next()
	}

	intConversion, _ := strconv.Atoi(result)

	return intConversion
}

func (t *Tokenizer) compareString(value string) bool {
	return strings.Compare(string(*t.CurrentChar), value) == 0
}

func (t *Tokenizer) GetToken() Token {
	if t.Pos == 0 {
		t.CurrentChar = new(byte)
		*t.CurrentChar = t.Text[t.Pos]
	}

	for t.CurrentChar != nil {

		if unicode.IsSpace(rune(*t.CurrentChar)) {
			t.Next()
			continue
		}

		if regexp.MustCompile(`\d`).Match([]byte{*t.CurrentChar}) {
			return Token{INTEGER, t.digit()}
		}

		if t.compareString("(") || t.compareString(")") {
			a := *t.CurrentChar
			t.Next()
			return Token{PARENTHESES, string(a)}
		}

		if t.compareString("+") {
			t.Next()
			return Token{PLUS, "+"}
		}

		if t.compareString("-") {
			t.Next()
			return Token{MINUS, "-"}
		}

		if t.compareString("*") {
			t.Next()
			return Token{MULTIPLICATION, "*"}
		}

		if t.compareString("/") {
			t.Next()
			return Token{DIVISION, "/"}
		}

		panic("Parsing Next token")
	}

	return Token{EOF, ""}
}
