package ast

import (
	"github.com/falcinspire/scriptblock/front/symbols"
)

// Expression is any node that represents an expression
type Expression interface {
	Accept(visitor ExpressionVisitor)
}

// NumberExpression is an expression that represents a number
type NumberExpression struct {
	Value float64
}

// NewNumberExpression is a constructor for NumberExpression
func NewNumberExpression(value float64) *NumberExpression {
	expr := new(NumberExpression)
	expr.Value = value
	return expr
}

// Accept runs the double dispatch for the visitor
func (expression *NumberExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitNumber(expression)
}

// AddExpression is any node that represents a addition expression
type AddExpression struct {
	Left  Expression
	Right Expression
}

// NewAddExpression is a constructor for AddExpression
func NewAddExpression(left Expression, right Expression) *AddExpression {
	expression := new(AddExpression)
	expression.Left = left
	expression.Right = right
	return expression
}

// Accept runs the double dispatch for the visitor
func (expression *AddExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitAdd(expression)
}

// SubtractExpression is any node that represents a subtraction expression
type SubtractExpression struct {
	Left  Expression
	Right Expression
}

// NewSubtractExpression is a constructor for SubtractExpression
func NewSubtractExpression(left Expression, right Expression) *SubtractExpression {
	expression := new(SubtractExpression)
	expression.Left = left
	expression.Right = right
	return expression
}

// Accept runs the double dispatch for the visitor
func (expression *SubtractExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitSubtract(expression)
}

// MultiplyExpression is any node that represents a multiply expression
type MultiplyExpression struct {
	Left  Expression
	Right Expression
}

// NewMultiplyExpression is a constructor for MultiplyExpression
func NewMultiplyExpression(left Expression, right Expression) *MultiplyExpression {
	expression := new(MultiplyExpression)
	expression.Left = left
	expression.Right = right
	return expression
}

// Accept runs the double dispatch for the visitor
func (expression *MultiplyExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitMultiply(expression)
}

// DivideExpression is any node that represents a divide expression
type DivideExpression struct {
	Left  Expression
	Right Expression
}

// NewDivideExpression is a constructor for DivideExpression
func NewDivideExpression(left Expression, right Expression) *DivideExpression {
	expression := new(DivideExpression)
	expression.Left = left
	expression.Right = right
	return expression
}

// Accept runs the double dispatch for the visitor
func (expression *DivideExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitDivide(expression)
}

// IntegerDivideExpression is any node that represents a integer divide expression
type IntegerDivideExpression struct {
	Left  Expression
	Right Expression
}

// NewIntegerDivideExpression is a constructor for DivideExpression
func NewIntegerDivideExpression(left Expression, right Expression) *IntegerDivideExpression {
	expression := new(IntegerDivideExpression)
	expression.Left = left
	expression.Right = right
	return expression
}

// Accept runs the double dispatch for the visitor
func (expression *IntegerDivideExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitIntegerDivide(expression)
}

// PowerExpression is any node that represents a power expression
type PowerExpression struct {
	Left  Expression
	Right Expression
}

// NewPowerExpression is a constructor for PowerExpression
func NewPowerExpression(left Expression, right Expression) *PowerExpression {
	expression := new(PowerExpression)
	expression.Left = left
	expression.Right = right
	return expression
}

// Accept runs the double dispatch for the visitor
func (expression *PowerExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitPower(expression)
}

// StringExpression is any node that represents a string expression
type StringExpression struct {
	Value string
}

// NewStringExpression is a constructor for StringExpression
func NewStringExpression(value string) *StringExpression {
	expr := new(StringExpression)
	expr.Value = value
	return expr
}

// Accept runs the double dispatch for the visitor
func (expression *StringExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitString(expression)
}

// FormatterExpression is any node that represents a formatter expression
type FormatterExpression struct {
	Format    string
	Arguments []Expression
}

// NewFormatterExpression is a constructor for FormatterExpression
func NewFormatterExpression(format string, arguments []Expression) *FormatterExpression {
	formatter := new(FormatterExpression)
	formatter.Format = format
	formatter.Arguments = arguments
	return formatter
}

// Accept runs the double dispatch for the visitor
func (expression *FormatterExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitFormatter(expression)
}

// IdentifierExpression is any node that represents a identifier expression
type IdentifierExpression struct {
	Name    string
	Address *symbols.AddressBox // resolution pass
	Free    bool                // resolution pass
}

// NewIdentifierExpression is a constructor for IdentifierExpression
func NewIdentifierExpression(name string) *IdentifierExpression {
	return &IdentifierExpression{name, nil, false}
}

// Accept runs the double dispatch for the visitor
func (expression *IdentifierExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitIdentifier(expression)
}

// ClosureExpression is any node that represents a closure expression
type ClosureExpression struct {
	Callee  Expression
	Capture []*IdentifierExpression
}

// NewClosureExpression is a constructor for ClosureExpression
func NewClosureExpression(callee Expression, capture []*IdentifierExpression) *ClosureExpression {
	return &ClosureExpression{callee, capture}
}

// Accept runs the double dispatch for the visitor
func (expression *ClosureExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitClosure(expression)
}

// CallExpression is any node that represents a call expression
type CallExpression struct {
	Identifier Expression
	Arguments  []Expression
	Captures   []Expression
}

// NewCallExpression is a constructor for CallExpression
func NewCallExpression(identifier Expression, arguments []Expression) *CallExpression {
	return &CallExpression{identifier, arguments, []Expression{}}
}

// Accept runs the double dispatch for the visitor
func (expression *CallExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitCall(expression)
}
