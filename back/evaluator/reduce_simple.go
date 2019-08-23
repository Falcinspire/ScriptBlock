package evaluator

import (
	"fmt"

	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/symbols"
)

func ReduceIdentifier(identifier *ast.IdentifierExpression, data *EvaluateData) values.Value {
	visitor := NewReduceExpressionVisitor(data)
	visitor.VisitIdentifier(identifier)
	return visitor.Result
}

func ReduceExpression(expression ast.Expression, data *EvaluateData) values.Value {
	return NewReduceExpressionVisitor(data).QuickVisitExpression(expression)
}

func ReduceArgumentList(arguments []ast.Expression, data *EvaluateData) []values.Value {
	argumentValues := make([]values.Value, len(arguments))
	expressionVisitor := NewReduceExpressionVisitor(data)
	for i, argument := range arguments {
		argumentValues[i] = expressionVisitor.QuickVisitExpression(argument)
	}
	return argumentValues
}

func GetValueForAddress(address *symbols.AddressBox, data *EvaluateData) values.Value {
	switch address.Type {
	case symbols.UNIT:
		unitAddress := address.Data.(*symbols.UnitAddress)
		if unitAddress.Module == data.Location.Module && unitAddress.Unit == data.Location.Unit {
			// this unit
			selfExported := values.LookupValueTable(unitAddress.Module, unitAddress.Unit, data.ValueLibrary.Exported)
			result, exists := values.LookupUnitValue(unitAddress.Module, unitAddress.Unit, unitAddress.Name, selfExported)
			if exists {
				return result
			}
			selfInternal := values.LookupValueTable(unitAddress.Module, unitAddress.Unit, data.ValueLibrary.Internal)
			internalResult, exists := values.LookupUnitValue(unitAddress.Module, unitAddress.Unit, unitAddress.Name, selfInternal)
			if exists {
				return internalResult
			}

			panic(fmt.Errorf("Could not get value for address %s", unitAddress))
		} else {
			// imported unit
			table := values.LookupValueTable(unitAddress.Module, unitAddress.Unit, data.ValueLibrary.Exported)
			result, exists := values.LookupUnitValue(unitAddress.Module, unitAddress.Unit, unitAddress.Name, table)
			if !exists {
				panic(fmt.Errorf("Could not get value for address %s", unitAddress))
			}
			return result
		}
	case symbols.PARAMETER:
		paramAddress := address.Data.(*symbols.ParameterAddress)
		result, exists := values.LookupParameterValue(paramAddress.Name, data.LocalValues)
		if !exists {
			panic(fmt.Errorf("Could not get value for address %s", paramAddress))
		}
		return result
	case symbols.CAPTURE:
		captureAddress := address.Data.(*symbols.CaptureAddress)
		result, exists := values.LookupCaptureValue(captureAddress.Name, data.LocalValues)
		if !exists {
			panic(fmt.Errorf("Could not get value for address %s", captureAddress))
		}
		return result
	}

	panic(fmt.Errorf("Fell out of switch? %s", address))
}
