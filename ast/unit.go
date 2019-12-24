package ast

// Unit is any node that represents a unit
type Unit struct {
	ImportLines []*ImportLine
	Definitions []TopDefinition
}

// NewUnit is a constructor for Unit
func NewUnit(importLines []*ImportLine, definitions []TopDefinition) *Unit {
	unit := new(Unit)
	unit.ImportLines = importLines
	unit.Definitions = definitions
	return unit
}

// Accept runs the double dispatch for the visitor
func (unit *Unit) Accept(visitor UnitVisitor) {
	visitor.VisitUnit(unit)
}
