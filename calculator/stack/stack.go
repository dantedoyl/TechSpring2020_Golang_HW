package stack

type Stack struct {
	data []interface{}
}

func NewStack() Stack {
	return Stack{
		data: []interface{}{},
	}
}

func (s Stack) IsEmpty() bool {
	return len((s).data) == 0
}

func (s *Stack) Push(el interface{}) {
	(*s).data = append((*s).data, el)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		ind := len((*s).data) - 1
		el := (*s).data[ind]
		(*s).data = (*s).data[:ind]
		return el, true
	}
}

func (s Stack) Top() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		ind := len((s).data) - 1
		el := (s).data[ind]
		return el, true
	}
}
