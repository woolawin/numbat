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
		stack.exprStack.Pop()
		return
	}

	stack.calls = stack.calls[:len(stack.calls)-1]
	stack.current = stack.calls[len(stack.calls)-1]
	stack.exprStack.Pop()
}

func (stack *CallStack) Current() *Call {
	return stack.current
}

type TypeStack struct {
	current *Type
	types   []*Type
}

func (stack *TypeStack) Push(t *Type) {
	stack.types = append(stack.types, t)
	stack.current = t
}

func (stack *TypeStack) Pop() {
	if len(stack.types) == 0 {
		return
	}

	if len(stack.types) == 1 {
		stack.current = stack.types[0]
		stack.types = nil
		return
	}

	stack.types = stack.types[:len(stack.types)-1]
	stack.current = stack.types[len(stack.types)-1]
}

func (stack *TypeStack) Current() *Type {
	return stack.current
}
