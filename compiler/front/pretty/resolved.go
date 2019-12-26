package pretty

type ResolvedUnit struct {
	Imports     []*ResolvedImport
	Definitions []ResolvedTopDefinition
}

type ResolvedImport struct {
	Module string
	Unit   string
}

type ResolvedTopDefinition interface{}

type ResolvedConstantDefinition struct {
	Type          string
	Name          string
	Internal      bool
	Documentation string
	Value         ResolvedExpression
}

type ResolvedFunctionDefinition struct {
	Type          string
	Name          string
	Internal      bool
	Documentation string
	Body          []ResolvedStatement
}

type ResolvedTemplateDefinition struct {
	Type          string
	Name          string
	Internal      bool
	Documentation string
	Parameters    []string
	Captures      []string
	Body          []ResolvedStatement
}

type ResolvedStatement interface{}

type ResolvedFunctionCall struct {
	Type      string
	Callee    ResolvedExpression
	Arguments []ResolvedExpression
	Trailing  *TrailingFunction
}

type TrailingFunction struct {
	Type       string
	Parameters []string
	Body       []ResolvedStatement
}

type ResolvedNativeCall struct {
	Type      string
	Arguments []ResolvedExpression
}

type ResolvedExpression interface{}

type ResolvedIdentifierExpression struct {
	Type    string
	Address ResolvedAddress
}

type ResolvedOperatorExpression struct {
	Operator    string
	Left, Right ResolvedExpression
}

type ResolvedNumberExpression struct {
	Type  string
	Value float64
}

type ResolvedStringExpression struct {
	Type  string
	Value string
}

type ResolvedAddress interface{}

type ResolvedUnitAddress struct {
	Module, Unit, Name string
}
type ResolvedParameterAddress struct {
	FunctorDepth  int
	ParameterName string
}
type ResolvedCaptureAddress struct {
	CaptureName string
}
