package ast

// ExpressionVisitor is a visitor for expressions
type ExpressionVisitor interface {
	// VisitNumber is the double dispatch for number expressions
	VisitNumber(expr *NumberExpression)
	// VisitString is the double dispatch for string expressions
	VisitString(expr *StringExpression)
	// VisitAdd is the double dispatch for addition expressions
	VisitAdd(expr *AddExpression)
	// VisitSubtract is the double dispatch for subtraction expressions
	VisitSubtract(expr *SubtractExpression)
	// VisitMultiply is the double dispatch for multiply expressions
	VisitMultiply(expr *MultiplyExpression)
	// VisitDivide is the double dispatch for divide expressions
	VisitDivide(expr *DivideExpression)
	// VisitIntegerDivide is the double dispatch for integer divide expressions
	VisitIntegerDivide(expr *IntegerDivideExpression)
	// VisitPower is the double dispatch for power expressions
	VisitPower(expr *PowerExpression)
	// VisitFormatter is the double dispatch for formatter expressions
	VisitFormatter(expr *FormatterExpression)
	// VisitIdentifier is the double dispatch for identifier expressions
	VisitIdentifier(expr *IdentifierExpression)
	// VisitFunctor is the double dispatch for closure expressions
	VisitFunctor(expr *FunctorExpression)
	// VisitCall is the double dispatch for call expressions
	VisitCall(expr *CallExpression)
}

// StatementVisitor is the visitor for statements
type StatementVisitor interface {
	// VisitFunctionCall is the double dispatch for function call statements
	VisitFunctionCall(functionCall *FunctionCall)
	// VisitNativeCall is the double dispatch for native call statements
	VisitNativeCall(nativeCall *NativeCall)
	// VisitDelay is the double dispatch for delay statements
	VisitDelay(delay *DelayStatement)
	// VisitRaise is the double dispatch for raise statements
	VisitRaise(delay *RaiseStatement)
}

// TopVisitor is the visitor for top definitions
type TopVisitor interface {
	// VisitConstantDefinition is the visitor for constant definitions
	VisitConstantDefinition(definition *ConstantDefinition)
	// VisitFunctionDefinition is the visitor for function definitions
	VisitFunctionDefinition(definition *FunctionDefinition)
	// VisitFunctionShortcutDefinition is the visitor for function shortcut definitions
	VisitFunctionShortcutDefinition(shortcut *FunctionShortcutDefinition)
	// VisitTemplateDefinition is the visitor for closure definitions
	VisitTemplateDefinition(definition *TemplateDefinition)
}

// UnitVisitor is the visitor for units
type UnitVisitor interface {
	VisitUnit(unit *Unit)
}
