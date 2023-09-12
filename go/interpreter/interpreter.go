package interpreter

import (
	"reflect"
)

type NodeVisitor interface {
	visit(node Expr)
}

type Interpreter struct {
	parser Parser
}

func (i *Interpreter) AddParser(parser Parser) {
	i.parser = parser
}

func (i *Interpreter) visit(node Expr) Expr {
	nodeType := reflect.TypeOf(node).Name()

	switch nodeType {
	case "BinaryOperator":
		return i.visitBinaryOperator(node.(BinaryOperator))
	case "Integer":
		return i.visitInteger(node.(Integer))
	default:
		return node
	}
}

func (i *Interpreter) visitBinaryOperator(node BinaryOperator) int {
	left, right := node.left, node.right

	switch node.op {
	case PLUS:
		return i.visit(left).(int) + i.visit(right).(int)
	case MINUS:
		return i.visit(left).(int) - i.visit(right).(int)
	case MULTIPLICATION:
		return i.visit(left).(int) * i.visit(right).(int)
	case DIVISION:
		return i.visit(left).(int) / i.visit(right).(int)
	}

	panic("WTF ARE YOU DOING HERE?")
}

func (i *Interpreter) visitInteger(node Integer) int {
	return node.Value()
}

func (i *Interpreter) Interpret() Expr {
	tree := i.parser.Parse()
	return i.visit(tree)
}
