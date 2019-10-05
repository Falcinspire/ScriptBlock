package ast

// ImportLine is any node that represents an import line
type ImportLine struct {
	Location string
	Module   string
	Unit     string
	Metadata *Metadata
}
