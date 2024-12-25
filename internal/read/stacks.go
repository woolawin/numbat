package read

type ExprGroupStack struct {
	current *ExprGroup
	values  []*ExprGroup
}

func (stack *ExprGroupStack) Push(grp *ExprGroup) {
	stack.values = append(stack.values, grp)
	stack.current = grp
}

func (stack *ExprGroupStack) Pop() {
	if len(stack.values) == 0 {
		return
	}

	if len(stack.values) == 1 {
		stack.current = nil
		stack.values = nil
		return
	}

	stack.values = stack.values[:len(stack.values)-1]
	stack.current = stack.values[len(stack.values)-1]
}

func (stack *ExprGroupStack) Current() *ExprGroup {
	return stack.current
}

type CallStack struct {
	current   *Call
	calls     []*Call
	exprStack *ExprGroupStack
}

func (stack *CallStack) Push(call *Call) {
	stack.calls = append(stack.calls, call)
	stack.current = call

	stack.exprStack.Push(&stack.current.Arguments)
}

func (stack *CallStack) Pop() {
	if len(stack.calls) == 0 {
		return
	}

	if len(stack.calls) == 1 {
		stack.current = nil
		stack.calls = nil
		return
	}

	stack.calls = stack.calls[:len(stack.calls)-1]
	stack.current = stack.calls[len(stack.calls)-1]
	stack.exprStack.Pop()
}

func (stack *CallStack) Current() *Call {
	return stack.current
}
