package mystack
import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func(stack Stack) Top() (interface{}, error) {
	if(len(stack) == 0){
		return nil, errors.New("stack is empty")
	}

	return stack[len(stack)-1], nil
}
