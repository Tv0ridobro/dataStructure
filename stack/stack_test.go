package stack

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestStack_Peek(t *testing.T) {
	t.Parallel()
	stack := New[int]()
	stack.Peek()
	stack.Push(0)
	stack.Peek()
	if v := stack.Peek(); v != 0 {
		t.Errorf("wrong peek %d != %d", v, 0)
	}
	stack.Push(1)
	if v := stack.Peek(); v != 1 {
		t.Errorf("wrong peek %d != %d", v, 1)
	}
	stack.Push(2)
	if v := stack.Peek(); v != 2 {
		t.Errorf("wrong peek %d != %d", v, 2)
	}
	stack.Push(3)
	if v := stack.Peek(); v != 3 {
		t.Errorf("wrong peek %d != %d", v, 3)
	}
	stack.Push(4)
	if v := stack.Peek(); v != 4 {
		t.Errorf("wrong peek %d != %d", v, 4)
	}
	if v := stack.Peek(); v != 4 {
		t.Errorf("wrong second peek %d != %d", v, 4)
	}
	stack.Pop()
	if v := stack.Peek(); v != 3 {
		t.Errorf("wrong peek %d != %d", v, 3)
	}
}

func TestStack_Pop(t *testing.T) {
	t.Parallel()
	stack := New[int]()
	if v := stack.Pop(); v != 0 {
		t.Errorf("wrong pop %d != %d", v, 0)
	}
	stack.Push(90)
	if v := stack.Pop(); v != 90 {
		t.Errorf("wrong pop %d != %d", v, 90)
	}
	stack.Push(99)
	if v := stack.Pop(); v != 99 {
		t.Errorf("wrong pop %d != %d", v, 99)
	}
}

func TestStack_Size(t *testing.T) {
	t.Parallel()
	stack := New[int]()
	if stack.Size() != 0 {
		t.Errorf("wrong size of stack")
	}
	stack.Push(0)
	if stack.Size() != 1 {
		t.Errorf("wrong size of stack")
	}
	stack.Push(1)
	if stack.Size() != 2 {
		t.Errorf("wrong size of stack")
	}
	stack.Push(0)
	if stack.Size() != 3 {
		t.Errorf("wrong size of stack")
	}
	stack.Push(1)
	if stack.Size() != 4 {
		t.Errorf("wrong size of stack")
	}
	stack.Push(0)
	if stack.Size() != 5 {
		t.Errorf("wrong size of stack")
	}
	stack.Pop()
	if stack.Size() != 4 {
		t.Errorf("wrong size of stack")
	}
	stack.Push(1)
	if stack.Size() != 5 {
		t.Errorf("wrong size of stack")
	}
	stack.Pop()
	if stack.Size() != 4 {
		t.Errorf("wrong size of stack")
	}
	stack.Pop()
	if stack.Size() != 3 {
		t.Errorf("wrong size of stack")
	}
}

func TestStack_All(t *testing.T) {
	t.Parallel()
	stack := New[int]()
	if !slices.Equal(stack.All(), []int{}) {
		t.Errorf("wrong all %v", []int{})
	}
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	if !slices.Equal(stack.All(), []int{0, 1, 2, 3, 4}) {
		t.Errorf("wrong all %v", []int{0, 1, 2, 3, 4})
	}
	stack.Pop()
	if !slices.Equal(stack.All(), []int{0, 1, 2, 3}) {
		t.Errorf("wrong all %v", []int{0, 1, 2, 3})
	}
	stack.Push(7)
	if !slices.Equal(stack.All(), []int{0, 1, 2, 3, 7}) {
		t.Errorf("wrong all %v", []int{0, 1, 2, 3, 7})
	}
	if !slices.Equal(stack.All(), []int{0, 1, 2, 3, 7}) {
		t.Errorf("wrong all %v", []int{0, 1, 2, 3, 7})
	}
}
