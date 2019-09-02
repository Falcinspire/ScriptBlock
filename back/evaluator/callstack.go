package evaluator

// TODO need function location
type CallStack struct {
	stack []string
}

func PushCallStack(function string, stack *CallStack) {
	stack.stack = append(stack.stack, function)
}

func PopCallStack(stack *CallStack) {
	stack.stack = stack.stack[0 : len(stack.stack)-1]
}

func ListCallStack(stack *CallStack) []string {
	return stack.stack
}

func NewCallStack() *CallStack {
	return &CallStack{[]string{}}
}
