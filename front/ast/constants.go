package ast

// ConstantDefinition is a node representing a top level definition of a constant
type ConstantDefinition struct {
	Name          string
	Value         Expression
	Internal      bool
	Documentation string
	Metadata      *Metadata
}

// NewConstantDefinition is a constructor for ConstantDefinition
func NewConstantDefinition(name string, value Expression, internal bool, docs string, metadata *Metadata) *ConstantDefinition {
	return &ConstantDefinition{name, value, internal, docs, metadata}
}

// Accept runs the double dispatch for the visitor
func (definition *ConstantDefinition) Accept(visitor TopVisitor) {
	visitor.VisitConstantDefinition(definition)
}
