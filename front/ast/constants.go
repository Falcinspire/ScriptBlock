package ast

// ConstantDefinition is a node representing a top level definition of a constant
type ConstantDefinition struct {
	Name          string
	Value         Expression
	Internal      bool
	Documentation string
}

// NewConstantDefinition is a constructor for ConstantDefinition
func NewConstantDefinition(name string, value Expression, internal bool, docs string) *ConstantDefinition {
	return &ConstantDefinition{name, value, internal, docs}
}

// Accept runs the double dispatch for the visitor
func (definition *ConstantDefinition) Accept(visitor TopVisitor) {
	visitor.VisitConstantDefinition(definition)
}
