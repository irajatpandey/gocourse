package main

import "fmt"

// Stack struct represents a stack data structure
type Stack struct {
    top      int   // Index of the top element in the stack
    capacity int   // Current capacity of the stack
    data     []int // Underlying slice to store stack elements
}

// initializeStack creates and returns a new stack with initial capacity
func initializeStack() *Stack {
    capacity := 2
    return &Stack{
        top:      -1,
        capacity: capacity,
        data:     make([]int, capacity),
    }
}

// isEmpty checks if the stack is empty
func (st *Stack) isEmpty() bool {
    return st.top == -1
}

// isFull checks if the stack is full
func (st *Stack) isFull() bool {
    return st.top == st.capacity-1
}

// resize doubles the capacity of the stack when it is full
func (st *Stack) resize() {
    newCapacity := 2 * st.capacity
    newData := make([]int, newCapacity)

    // Copy old data to new array
    for i := 0; i < st.capacity; i++ {
        newData[i] = st.data[i]
    }

    st.capacity = newCapacity
    st.data = newData
}

// push adds a new element to the top of the stack
func (st *Stack) push(value int) {
    if st.isFull() {
        st.resize()
    }
    st.top++
    st.data[st.top] = value
}

// pop removes and returns the top element from the stack
func (st *Stack) pop() int {
    if st.isEmpty() {
        fmt.Println("Stack is empty...!")
        return -1
    }

    st.top--
    poppedElement := st.data[st.top+1]
    return poppedElement
}

func main() {
    // Initialize a new stack
    stack := initializeStack()
    fmt.Println(stack.data) // Print initial stack data (all zeros)

    // Push elements onto the stack
    stack.push(10)
    stack.push(20)
    stack.push(30)
    stack.push(40)
    stack.push(50)

    fmt.Println(stack.data) // Print stack data after pushes

    // Pop elements from the stack
    fmt.Println("Popped Element", stack.pop())
    fmt.Println("Popped Element", stack.pop())

    // Push another element
    stack.push(60)
    fmt.Println(stack.data) // Print stack data after another push
}