package evaluator

import "github.com/falcinspire/scriptblock/back/values"

func RawifyValue(value values.Value) string {
	visitor := NewRawifyValueVisitor()
	value.Accept(visitor)
	return visitor.Result
}
