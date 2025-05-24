package main

import "fmt"

// ArrayList is a dynamic array implementation in Go
type ArrayList struct {
    size     int   // Number of elements currently in the array
    capacity int   // Current capacity of the underlying array
    data     []int // Underlying slice to store elements
}

// newDynamicArray creates and returns a new ArrayList with initial capacity
func newDynamicArray() *ArrayList {
    initialCapacity := 2
    return &ArrayList{
        data:     make([]int, initialCapacity),
        size:     0,
        capacity: initialCapacity,
    }
}

// resize doubles the capacity of the ArrayList when it is full
func (ls *ArrayList) resize() {
    newCapacity := ls.capacity * 2
    newData := make([]int, newCapacity)

    // Copy old data to new array
    for i := 0; i < ls.capacity; i++ {
        newData[i] = ls.data[i]
    }
    ls.data = newData
    ls.capacity = newCapacity
}

// insert adds a new element to the end of the ArrayList
func (ls *ArrayList) insert(value int) {
    if ls.size == ls.capacity {
        ls.resize()
    }
    ls.data[ls.size] = value
    ls.size++
}

// printArray prints all elements currently in the ArrayList
func printArray(arr *ArrayList) {
    for i := 0; i < arr.size; i++ {
        fmt.Print(arr.data[i], " ")
    }
    fmt.Println()
}

func main() {
    fmt.Println("ArrayList in Go")

    // Create a new dynamic array
    arr := newDynamicArray()

    // Insert elements
    arr.insert(10)
    arr.insert(20)
    arr.insert(30)

    // Print the array
    printArray(arr)
}