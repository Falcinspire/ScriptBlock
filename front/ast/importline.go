package ast

// ImportLine is any node that represents an import line
type ImportLine struct {
	Module string
	Unit   string
}
