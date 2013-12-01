package queue

// Simple queue data structure using the standard List container

import ("fmt"
        "container/list")

type Queue struct {
    list list.List
}

// Initializes a new Queue. Returns *Queue
func NewQueue() *Queue {
    return new(Queue)
}

// Returns the Value of the first element in the queue
func (q Queue) Peek() interface{} {
    front := q.list.Front()
    return front.Value
}

// Pushes an item to the back of the queue
func (q *Queue) Push(item interface{}) {
    q.list.PushBack(item)
}

// Pops an item from the back of the queue
func (q *Queue) Pop() interface{} {
    fmt.Printf("Pop function\n")
    var ele = q.list.Remove(q.list.Front())
    return ele
}

// Lists all the elements in the queue in order
func (q Queue) Print() {
    for element := q.list.Front(); element != nil; element = element.Next() {
        fmt.Printf("%v\n", element.Value)
    }
}

// Returns the size of the queue
func (q Queue) Size() int {
    return q.list.Len()
}

// Checks if the queue is empty or not
func (q Queue) IsEmpty() bool {
    return q.list.Len() == 0
}
