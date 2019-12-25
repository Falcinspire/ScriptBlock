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

type FunctorValue struct {
	Callee  *TemplateValue
	Capture []Value
}

func NewFunctorValue(callee *TemplateValue, capture []Value) *FunctorValue {
	return &FunctorValue{callee, capture}
}
func (this *FunctorValue) Accept(visitor ValueVisitor) {
	visitor.VisitFunctor(this)
}
