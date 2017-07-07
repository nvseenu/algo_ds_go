package collections

import (
	"testing"
)

func TestStackEmpty(t *testing.T) {

	stack, _ := NewStack()
	if !stack.IsEmpty() {
		t.Error("Stack is not empty when it has no items")
	}
}

func TestStackPush(t *testing.T) {

	stack, _ := NewStack()
	stack.Push(10)
	if stack.Size() != 1 {
		t.Errorf("Expected Stack size is 1 but got %v", stack.Size())
	}
	stack.Push(11)
	if stack.Size() != 2 {
		t.Errorf("Expected Stack size is 2 but got %v", stack.Size())
	}
	stack.Push(12)
	if stack.Size() != 3 {
		t.Errorf("Expected Stack size is 3 but got %v", stack.Size())
	}
}

func TestStackPushWhenStackIsFull(t *testing.T) {
	stack, _ := NewStack()
	for i := 1; i <= 16; i++ {
		stack.Push(i)
	}
	err := stack.Push(17)
	if err == nil {
		t.Error("Expected is error message, but got ", err)
	}
}

func TestStackPop(t *testing.T) {

	stack, _ := NewStack()
	stack.Push(10)
	stack.Push(11)
	stack.Push(12)

	elm, _ := stack.Pop()
	eval, _ := elm.(int)
	if eval != 12 {
		t.Errorf("Expected 12 for the first popup, but got %v", eval)
	}

	if stack.Size() != 2 {
		t.Errorf("Expected Stack size 2 after the first popup, but got %v", stack.Size())
	}

	elm, _ = stack.Pop()
	eval, _ = elm.(int)
	if eval != 11 {
		t.Errorf("Expected 11 for the second popup, but got %v", eval)
	}

	if stack.Size() != 1 {
		t.Errorf("Expected Stack size 1 after the second popup, but got %v", stack.Size())
	}

	elm, _ = stack.Pop()
	eval, _ = elm.(int)
	if eval != 10 {
		t.Errorf("Expected 10 for the third popup, but got %v", eval)
	}

	if stack.Size() != 0 {
		t.Errorf("Expected Stack size 0 after the thrird popup, but got %v", stack.Size())
	}

}

func TestStackPopWhenStackIsEmpty(t *testing.T) {

	stack, _ := NewStack()

	_, err := stack.Pop()

	if err == nil {
		t.Error("Expected error message, but got ", err)
	}
}

func TestStackPeek(t *testing.T) {

	stack, _ := NewStack()
	stack.Push(10)
	stack.Push(11)
	stack.Push(12)

	elm, _ := stack.Peek()
	eval, _ := elm.(int)
	if eval != 12 {
		t.Errorf("Expected 12 for peek, but got %v", eval)
	}

	if stack.Size() != 3 {
		t.Errorf("Expected same stack size 3 after the peek, but got %v", stack.Size())
	}

}

func TestStackPeekWhenStackIsEmpty(t *testing.T) {

	stack, _ := NewStack()

	_, err := stack.Peek()
	if err == nil {
		t.Errorf("Expected error message, but got nil")
	}
}
