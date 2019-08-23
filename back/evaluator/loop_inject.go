package evaluator

import (
	"fmt"

	"github.com/falcinspire/scriptblock/back/dumper"
	"github.com/falcinspire/scriptblock/back/output"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/google/uuid"
)

type LoopInject struct {
	InjectBody []string
}

func NewLoopInject() *LoopInject {
	return &LoopInject{[]string{}}
}

func GenerateTickFunction(inject *LoopInject, unitLocation *location.UnitLocation, output output.OutputDirectory) (module, unit, name string) {
	loopFunctionName := fmt.Sprintf("loop-%s", uuid.New().String())
	loopFunctionBody := inject.InjectBody
	dumper.DumpFunction(unitLocation.Module, unitLocation.Unit, loopFunctionName, loopFunctionBody, output)
	return unitLocation.Module, unitLocation.Unit, loopFunctionName
}

func GenerateDelayLines(inject *LoopInject, delay int, functionCall *ast.FunctionCall, data *EvaluateData) (runCommand string) {
	cloudID := uuid.New().String()
	invokeValue := ReduceExpression(ast.NewCallExpression(functionCall.Callee, functionCall.Arguments, nil), data)
	translateValue := RawifyValue(invokeValue)
	testCloud := createRepeatCommand(cloudID, delay, translateValue)
	inject.InjectBody = append(inject.InjectBody, testCloud)
	summonCloud := createSummonTimerCommand(cloudID, delay)
	return summonCloud
}

// createRepeatCommand generates the command needed to be looped for the delay to work
func createRepeatCommand(cloudID string, delay int, mcrunnable string) string {
	return fmt.Sprintf("execute as @e[type=minecraft:area_effect_cloud,tag=%s,nbt={Age:%d}] run %s", cloudID, delay, mcrunnable)
}

// createSummonTimerCommand generates the command needed to summon the cloud
func createSummonTimerCommand(cloudID string, delay int) string {
	return fmt.Sprintf("execute at @p run summon minecraft:area_effect_cloud ~ ~ ~ {Tags:[\"%s\"],Duration:%d}", cloudID, delay+1)
}
