package stack

import "testing"

func TestStack_IsEmpty(t *testing.T) {
	var stack Stack

	if !stack.IsEmpty() {
		t.Error("Expected empty stack.")
	}
}

func TestStack_Push(t *testing.T) {
	var stack Stack

	stack.Push("12")

	if exp, _ := stack.Pop(); exp != "12" {
		t.Error("Expected value.")
	}
}

func TestStack_Top(t *testing.T) {
	var stack Stack

	stack.Push("12")

	if exp, _ := stack.Top(); exp != "12" {
		t.Error("Expected value.")
	}

	stack.Pop()

	if _, ok := stack.Top(); ok {
		t.Error("Expected error.")
	}
}

func TestStack_Pop(t *testing.T) {
	var stack Stack

	if _, ok := stack.Pop(); ok {
		t.Error("Expected error.")
	}
}
