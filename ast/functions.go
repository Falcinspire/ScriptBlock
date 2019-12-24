package ast

// FunctionDefinition is a node representing a top level definition of a function
type FunctionDefinition struct {
	Name          string
	Body          []Statement
	Internal      bool
	Tag           Tag
	Documentation string
	Metadata      *Metadata
}

// NewFunctionDefinition is a constructor for FunctionDefinition
func NewFunctionDefinition(name string, body []Statement, internal bool, tag Tag, docs string, metadata *Metadata) *FunctionDefinition {
	return &FunctionDefinition{name, body, internal, tag, docs, metadata}
}

// Accept runs the double dispatch for the visitor
func (definition *FunctionDefinition) Accept(visitor TopVisitor) {
	visitor.VisitFunctionDefinition(definition)
}

// FunctionShortcutDefinition is a node representing a top level definition of a function shortcut
type FunctionShortcutDefinition struct {
	// shortcut name
	Name string

	// shortcut for
	FunctionCall *FunctionCall

	Internal      bool
	Tag           Tag
	Documentation string

	Metadata *Metadata
}

// NewFunctionShortcutDefinition is a constructor for FunctionShortcutDefinition
func NewFunctionShortcutDefinition(name string, functionCall *FunctionCall, internal bool, tag Tag, docs string, metadata *Metadata) *FunctionShortcutDefinition {
	return &FunctionShortcutDefinition{name, functionCall, internal, tag, docs, metadata}
}

// Accept runs the double dispatch for the visitor
func (definition *FunctionShortcutDefinition) Accept(visitor TopVisitor) {
	visitor.VisitFunctionShortcutDefinition(definition)
}

// TemplateDefinition is a node representing a top level definition of a template
type TemplateDefinition struct {
	Name          string
	Parameters    []string
	Body          []Statement
	Internal      bool
	Documentation string
	Metadata      *Metadata
}

// NewTemplateDefinition is a constructor for TemplateDefinition
func NewTemplateDefinition(name string, parameters []string, body []Statement, internal bool, docs string, metadata *Metadata) *TemplateDefinition {
	return &TemplateDefinition{name, parameters, body, internal, docs, metadata}
}

// Accept runs the double dispatch for the visitor
func (definition *TemplateDefinition) Accept(visitor TopVisitor) {
	visitor.VisitTemplateDefinition(definition)
}

// ClosureDefinition is a node representing an (injected) top level definition of a closure
type ClosureDefinition struct {
	Name       string
	Parameters []string
	Capture    []string
	Body       []Statement
	Internal   bool
	Metadata   *Metadata
}

// NewClosureDefinition is a constructor for ClosureDefinition
func NewClosureDefinition(name string, parameters []string, capture []string, body []Statement, internal bool, metadata *Metadata) *ClosureDefinition {
	return &ClosureDefinition{name, parameters, capture, body, internal, metadata}
}

// Accept runs the double dispatch for the visitor
func (definition *ClosureDefinition) Accept(visitor TopVisitor) {
	visitor.VisitClosureDefinition(definition)
}

// Tag is a node representing the tag for a function
type Tag struct {
	Namespace string
	Identity  string
	Metadata  *Metadata
}
