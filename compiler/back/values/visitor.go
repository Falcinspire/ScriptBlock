package values

type ValueVisitor interface {
	VisitNumber(numberValue *NumberValue)
	VisitString(stringValue *StringValue)
	VisitFunction(functionValue *FunctionValue)
	VisitTemplate(templateReference *TemplateValue)
	VisitFunctor(functorValue *FunctorValue)
}
