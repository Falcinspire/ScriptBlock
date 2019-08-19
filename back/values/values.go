package values

import (
	"fmt"
)

type Value interface {
	Accept(visitor ValueVisitor)
}

type NumberValue struct {
	Value float64
}

func NewNumberValue(value float64) *NumberValue {
	box := new(NumberValue)
	box.Value = value
	return box
}
func (this *NumberValue) Accept(visitor ValueVisitor) {
	visitor.VisitNumber(this)
}

func (this *NumberValue) String() string {
	return fmt.Sprintf("%.2f", this.Value)
}

type StringValue struct {
	Value string
}

func NewStringValue(value string) *StringValue {
	stringValue := new(StringValue)
	stringValue.Value = value
	return stringValue
}
func (this *StringValue) Accept(visitor ValueVisitor) {
	visitor.VisitString(this)
}

type FunctionValue struct {
	Module string
	Unit   string
	Name   string
}

func NewFunctionValue(module string, unit string, name string) *FunctionValue {
	return &FunctionValue{module, unit, name}
}
func (this *FunctionValue) Accept(visitor ValueVisitor) {
	visitor.VisitFunction(this)
}

type TemplateValue struct {
	Module string
	Unit   string
	Name   string
}

func NewTemplateValue(module string, unit string, name string) *TemplateValue {
	return &TemplateValue{module, unit, name}
}
func (this *TemplateValue) Accept(visitor ValueVisitor) {
	visitor.VisitTemplate(this)
}

// TODO rename function/template to function/template reference
type ClosureReferenceValue struct {
	Module, Unit, Name string
}

func NewClosureReferenceValue(module, unit, name string) *ClosureReferenceValue {
	return &ClosureReferenceValue{module, unit, name}
}
func (this *ClosureReferenceValue) Accept(visitor ValueVisitor) {
	visitor.VisitClosureReference(this)
}

type ClosureValue struct {
	Callee  *ClosureReferenceValue
	Capture []Value
}

func NewClosureValue(callee *ClosureReferenceValue, capture []Value) *ClosureValue {
	return &ClosureValue{callee, capture}
}
func (this *ClosureValue) Accept(visitor ValueVisitor) {
	visitor.VisitClosure(this)
}
