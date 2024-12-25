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
