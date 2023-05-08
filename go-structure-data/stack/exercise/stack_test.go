package exercise

import "testing"

func TestStackExercise(t *testing.T) {
	var stack *StackValue = NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for i := 5; i > 0; i-- {
		item := stack.Pop()

		if item != i {
			t.Error("Expected", i, "got", item)
		}
	}
}
