package ast

// TopDefinition is any node that represents a top definitino
type TopDefinition interface {
	Accept(visitor TopVisitor)
}
