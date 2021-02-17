package stack

type Stack struct {
	data []interface{}
}

func (stack *Stack) IsEmpty() bool {
	return len((*stack).data) == 0
}

func (stack *Stack) Push(el interface{}) {
	(*stack).data = append((*stack).data, el)
}

func (stack *Stack) Pop() (interface{}, bool) {
	if stack.IsEmpty() {
		return "", false
	} else {
		ind := len((*stack).data) - 1
		el := (*stack).data[ind]
		(*stack).data = (*stack).data[:ind]
		return el, true
	}
}

func (stack *Stack) Top() (interface{}, bool) {
	if stack.IsEmpty() {
		return "", false
	} else {
		ind := len((*stack).data) - 1
		el := (*stack).data[ind]
		return el, true
	}
}
