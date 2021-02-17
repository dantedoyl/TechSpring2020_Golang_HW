package calculate

import (
	"calculator/stack"
	"errors"
	"strconv"
)

func isNumb(ch string) bool {
	if ch != "+" && ch != "-" && ch != "/" && ch != "*" {
		return true
	} else {
		return false
	}
}

func getNextSymb(str *string) (string, error) {
	symb := ""
	if ch := (*str)[0]; ch >= 48 && ch <= 57 {
		for ; ch >= 48 && ch <= 57; ch = (*str)[0] {
			symb += string(ch)
			*str = (*str)[1:]
			if len(*str) == 0 {
				break
			}
		}
	} else {
		if ch != 40 && ch != 41 && ch != 42 && ch != 43 && ch != 45 && ch != 47 {
			return "", errors.New("Incorrect symbols in expression.")
		} else {
			symb += string(ch)
			*str = (*str)[1:]
		}
	}
	return symb, nil
}

func toPolishNotation(str *string) ([]string, error) {
	var polishStr []string
	var s stack.Stack
	for len(*str) > 0 {
		op, err := getNextSymb(str)
		if err != nil {
			return nil, err
		}
		switch op {
		case "(":
			s.Push(op)
		case ")":
			for op, ok := s.Top(); op != "(" && ok; op, ok = s.Top() {
				operand, _ := s.Pop()
				polishStr = append(polishStr, operand.(string))
			}
			_, ok := s.Top()
			if !ok {
				return nil, errors.New("In the expression either the delimiter is incorrectly supplied, or the parentheses are not matched")
			} else {
				s.Pop()
			}
		case "+", "-":
			if el, ok := s.Top(); ok && (el == "*") || (el == "/") {
				for ; ok && (el == "*") || (el == "/"); el, ok = s.Top() {
					operand, _ := s.Pop()
					polishStr = append(polishStr, operand.(string))
				}
				s.Push(op)
			} else {
				s.Push(op)
			}
		case "/", "*":
			s.Push(op)
		default:
			polishStr = append(polishStr, op)
		}
	}
	for !s.IsEmpty() {
		operand, _ := s.Pop()
		if isNumb(operand.(string)){
			return nil, errors.New("In the expression either the delimiter is incorrectly supplied, or the parentheses are not matched")
		}
		polishStr = append(polishStr, operand.(string))
	}
	return polishStr, nil
}

func performCalc(op1 int, op2 int, operator string) (int) {
	switch operator {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	default:
		return op1 / op2
	}
}

func Calculate(str string) (int, error) {
	ops, err := toPolishNotation(&str)
	if err != nil{
		return 0, err
	}

	var s stack.Stack
	for len(ops) > 0 {
		for ch := isNumb(ops[0]); ch; ch = isNumb(ops[0]) {
			numb, _ := strconv.Atoi(ops[0])
			s.Push(numb)
			ops = ops[1:]
		}
		secondOp, ok := s.Pop()
		if !ok {
			return 0, errors.New("In the expression either the delimiter or operator is incorrectly supplied.")
		}
		firstOp, ok := s.Pop()
		if !ok {
			return 0, errors.New("In the expression either the delimiter or operator is incorrectly supplied.")
		}
		result := performCalc(firstOp.(int), secondOp.(int), ops[0])
		s.Push(result)
		ops = ops[1:]
	}
	finalResult, _ := s.Pop()
	return finalResult.(int), nil
}