package ast

// Statement is any node that represents a statement
type Statement interface {
	Accept(visitor StatementVisitor)
}

// TrailingFunction is any node that represents a trailing function
type TrailingFunction struct {
	Parameters []string
	Body       []Statement
	Metadata   *Metadata
}

// FunctionCall is any node that represents a function call
type FunctionCall struct {
	Callee    Expression // TODO this is nearly always identifier, but can also be a closure
	Arguments []Expression
	Trailing  *TrailingFunction
	Metadata  *Metadata
}

// NewFunctionCall is a constructor for FunctionCall
func NewFunctionCall(callee Expression, arguments []Expression, trailing *TrailingFunction, metadata *Metadata) *FunctionCall {
	return &FunctionCall{callee, arguments, trailing, metadata}
}

// Accept runs the double dispatch for the visitor
func (statement *FunctionCall) Accept(visitor StatementVisitor) {
	visitor.VisitFunctionCall(statement)
}

// NativeCall is any node that represents a native call
type NativeCall struct {
	Arguments []Expression
	Metadata  *Metadata
}

// NewNativeCall is a constructor for NativeCall
func NewNativeCall(arguments []Expression, metadata *Metadata) *NativeCall {
	return &NativeCall{arguments, metadata}
}

// Accept runs the double dispatch for the visitor
func (statement *NativeCall) Accept(visitor StatementVisitor) {
	visitor.VisitNativeCall(statement)
}

// DelayStatement is any node that represents a delay statement
type DelayStatement struct {
	TickDelay    Expression
	Body         []Statement
	Metadata     *Metadata
	FunctionCall *FunctionCall // for when it is resolved
}

// NewDelayStatement is a constructor for DelayStatement
func NewDelayStatement(tickDelay Expression, body []Statement, metadata *Metadata) *DelayStatement {
	return &DelayStatement{tickDelay, body, metadata, nil}
}

// Accept runs the double dispatch for the visitor
func (statement *DelayStatement) Accept(visitor StatementVisitor) {
	visitor.VisitDelay(statement)
}

type RaiseStatement struct {
	Tag      Tag
	Metadata *Metadata
}

func NewRaiseStatement(tag Tag, metadata *Metadata) *RaiseStatement {
	return &RaiseStatement{tag, metadata}
}

// Accept runs the double dispatch for the visitor
func (statement *RaiseStatement) Accept(visitor StatementVisitor) {
	visitor.VisitRaise(statement)
}
