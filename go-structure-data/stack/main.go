package stack

type StackItemValue struct {
	// item is the value of the stack
	item interface{}
	// next is the pointer to the next item in the stack
	next *StackItemValue
}

type Stack struct {
	// sp is the stack pointer
	sp *StackItemValue
	// depth is the number of items in the stack / depth is a length of the stack
	depth uint64
}

// New is a constructor for the Stack
func New() *Stack {
	var stack *Stack = new(Stack)

	stack.depth = 0
	return stack
}

// Push adds an item to the stack
func (stack *Stack) Push(item any) {
	stack.sp = &StackItemValue{item: item, next: stack.sp}
	stack.depth++
}

// Pop removes an item from the stack
func (stack *Stack) Pop() any {
	if stack.depth > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}

	return nil
}

// Peek returns the top item from the stack without removing it
func (stack *Stack) Peek() any {
	if stack.depth > 0 {
		return stack.sp.item
	}

	return nil
}

//func main() {
//	stack := New()
//
//	stack.Push(1)
//	stack.Push(2)
//	stack.Push(3)
//
//	////fmt.Println(stack.Peek())
//	//fmt.Println(stack.Pop())
//	//fmt.Println(stack.Pop())
//	//fmt.Println(stack.Pop())
//
//	for i := 5; i > 0; i-- {
//		item := stack.Pop()
//		if item != i {
//			fmt.Println("Error")
//		}
//	}
//}
