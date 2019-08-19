package ast

// FunctionDefinition is a node representing a top level definition of a function
type FunctionDefinition struct {
	Name          string
	Body          []Statement
	Internal      bool
	Tag           Tag
	Documentation string
}

// NewFunctionDefinition is a constructor for FunctionDefinition
func NewFunctionDefinition(name string, body []Statement, internal bool, tag Tag, docs string) *FunctionDefinition {
	definition := new(FunctionDefinition)
	definition.Name = name
	definition.Body = body
	definition.Internal = internal
	definition.Tag = tag
	definition.Documentation = docs
	return definition
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
}

// NewFunctionShortcutDefinition is a constructor for FunctionShortcutDefinition
func NewFunctionShortcutDefinition(name string, functionCall *FunctionCall, internal bool, tag Tag, docs string) *FunctionShortcutDefinition {
	definition := new(FunctionShortcutDefinition)
	definition.Name = name
	definition.FunctionCall = functionCall
	definition.Internal = internal
	definition.Tag = tag
	definition.Documentation = docs
	return definition
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
}

// NewTemplateDefinition is a constructor for TemplateDefinition
func NewTemplateDefinition(name string, parameters []string, body []Statement, internal bool, docs string) *TemplateDefinition {
	definition := new(TemplateDefinition)
	definition.Name = name
	definition.Parameters = parameters
	definition.Body = body
	definition.Internal = internal
	definition.Documentation = docs
	return definition
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
}

// NewClosureDefinition is a constructor for ClosureDefinition
func NewClosureDefinition(name string, parameters []string, capture []string, body []Statement, internal bool) *ClosureDefinition {
	return &ClosureDefinition{name, parameters, capture, body, internal}
}

// Accept runs the double dispatch for the visitor
func (definition *ClosureDefinition) Accept(visitor TopVisitor) {
	visitor.VisitClosureDefinition(definition)
}

// Tag is a node representing the tag for a function
type Tag struct {
	Namespace string
	Identity  string
}
