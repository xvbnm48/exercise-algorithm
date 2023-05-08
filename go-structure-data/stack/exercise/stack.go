package exercise

type ThisIsStack struct {
	Item any
	Next *ThisIsStack
}

type StackValue struct {
	StackPointer *ThisIsStack
	Depth        uint64
}

func NewStack() *StackValue {
	var stack *StackValue = new(StackValue)

	stack.Depth = 0

	return stack
}

func (stack *StackValue) Push(item any) {
	stack.StackPointer = &ThisIsStack{Item: item, Next: stack.StackPointer}
	stack.Depth++
}

func (stack *StackValue) Pop() any {
	if stack.Depth > 0 {
		item := stack.StackPointer.Item
		stack.StackPointer = stack.StackPointer.Next
		stack.Depth--
		return item
	}

	return nil
}

func (stack *StackValue) Peek() any {
	if stack.Depth > 0 {
		return stack.StackPointer.Item
	}

	return nil

}
