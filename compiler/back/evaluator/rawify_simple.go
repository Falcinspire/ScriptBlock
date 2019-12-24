package evaluator

import "github.com/falcinspire/scriptblock/compiler/back/values"

func RawifyValue(value values.Value) string {
	visitor := NewRawifyValueVisitor()
	value.Accept(visitor)
	return visitor.Result
}
