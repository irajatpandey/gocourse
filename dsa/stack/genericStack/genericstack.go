package main

import "fmt"

type Stack[T any] struct {
	capacity int
	top int
	data []T
}

func initialiseStack[T any]() Stack[T] {
	capacity := 2
	top := -1

	return Stack[T]{
		capacity: capacity,
		top:      top,
		data:     make([]T, capacity),
	}	
}


func (st *Stack[T]) isEmpty() bool {
	if st.top == -1 {
		return true
	}

	return false
}

func (st *Stack[T]) push(element T) {
	st.top++
	if st.top >= st.capacity {
		st.capacity *= 2
		newData := make([]T, st.capacity)
		copy(newData, st.data)
		st.data = newData
	}
	st.data[st.top] = element
}

func (st *Stack[T]) pop() T {
	if st.isEmpty() {
		fmt.Println("Stack is empty...!")
		return *new(T) // Return zero value of type T
	}
	poppedElement := st.data[st.top]
	st.top--
	return poppedElement
}

func main() {

	stack := initialiseStack[int]()
	fmt.Println(stack.data)

	stack.push(10)
	stack.push(20)
	fmt.Println(stack.data) // Print stack data after pushing elements
	fmt.Println("Popped element:", stack.pop()) // Pop an element
	fmt.Println("Popped element:", stack.pop()) // Pop another element
	fmt.Println("Popped element:", stack.pop()) // Attempt to pop from an empty stack


}