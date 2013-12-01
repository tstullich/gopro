package stack

// Simple stack data structure using the standard List container

import ("fmt"
        "container/list")

type Stack struct {
    list list.List
}

// Initializes a new Stack. Returns *Stack
func NewStack() *Stack {
    return new(Stack)
}

// Returns the Value of the first element in the queue
func (s Stack) Peek() interface{} {
    front := s.list.Back()
    return front.Value
}

// Pushes an item to the back of the queue
func (s *Stack) Push(item interface{}) {
    s.list.PushBack(item)
}

// Pops an item from the back of the queue
func (s *Stack) Pop() interface{} {
    var ele = s.list.Remove(s.list.Back())
    return ele
}

// Lists all the elements in the queue in order
func (s Stack) Print() {
    for element := s.list.Back(); element != nil; element = element.Prev() {
        fmt.Printf("%v\n", element.Value)
    }
}

// Returns the size of the queue
func (s Stack) Size() int {
    return s.list.Len()
}

// Checks if the Stack is empty or not
func (s Stack) IsEmpty() bool {
    return s.list.Len() == 0
}
