package values

type BaseValueVisitor struct{}

func (visitor *BaseValueVisitor) VisitNumber(numberValue *NumberValue)           {}
func (visitor *BaseValueVisitor) VisitString(stringValue *StringValue)           {}
func (visitor *BaseValueVisitor) VisitFunction(functionValue *FunctionValue)     {}
func (visitor *BaseValueVisitor) VisitTemplate(templateReference *TemplateValue) {}
func (visitor *BaseValueVisitor) VisitClosure(closureValue *ClosureValue)        {}
