package ast

// BaseExpressionVisitor is a visitor for expressions that does nothing by default
type BaseExpressionVisitor struct{}

// VisitNumber is the double dispatch for number expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitNumber(expr *NumberExpression) {}

// VisitString is the double dispatch for string expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitString(expr *StringExpression) {}

// VisitAdd is the double dispatch for addition expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitAdd(expr *AddExpression) {}

// VisitSubtract is the double dispatch for subtraction expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitSubtract(expr *SubtractExpression) {}

// VisitMultiply is the double dispatch for multiply expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitMultiply(expr *MultiplyExpression) {}

// VisitDivide is the double dispatch for divide expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitDivide(expr *DivideExpression) {}

// VisitIntegerDivide is the double dispatch for integer divide expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitIntegerDivide(expr *IntegerDivideExpression) {}

// VisitPower is the double dispatch for power expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitPower(expr *PowerExpression) {}

// VisitFormatter is the double dispatch for formatter expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitFormatter(expr *FormatterExpression) {}

// VisitIdentifier is the double dispatch for identifier expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitIdentifier(expr *IdentifierExpression) {}

// VisitFunctor is the double dispatch for closure expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitFunctor(expr *FunctorExpression) {}

// VisitCall is the double dispatch for call expressions that does nothing by default
func (visitor *BaseExpressionVisitor) VisitCall(expr *CallExpression) {}

// BaseStatementVisitor is the visitor for statements that does nothing by default
type BaseStatementVisitor struct{}

// VisitFunctionCall is the double dispatch for function call statements that does nothing by default
func (visitor *BaseStatementVisitor) VisitFunctionCall(functionCall *FunctionCall) {}

// VisitNativeCall is the double dispatch for native call statements that does nothing by default
func (visitor *BaseStatementVisitor) VisitNativeCall(nativeCall *NativeCall) {}

// VisitDelay is the double dispatch for delay statements that does nothing by default
func (visitor *BaseStatementVisitor) VisitDelay(delay *DelayStatement) {}

// VisitRaise is the double dispatch for raise statements that does nothing by default
func (visitor *BaseStatementVisitor) VisitRaise(delay *RaiseStatement) {}

// BaseTopVisitor is the visitor for top definitions that does nothing by default
type BaseTopVisitor struct{}

// VisitConstantDefinition is the visitor for constant definitions that does nothing by default
func (visitor *BaseTopVisitor) VisitConstantDefinition(definition *ConstantDefinition) {}

// VisitFunctionDefinition is the visitor for function definitions that does nothing by default
func (visitor *BaseTopVisitor) VisitFunctionDefinition(definition *FunctionDefinition) {}

// VisitFunctionShortcutDefinition is the visitor for function shortcut definitions that does nothing by default
func (visitor *BaseTopVisitor) VisitFunctionShortcutDefinition(definition *FunctionShortcutDefinition) {
}

// VisitTemplateDefinition is the visitor for closure definitions that does nothing by default
func (visitor *BaseTopVisitor) VisitTemplateDefinition(definition *TemplateDefinition) {}
