package main

import "fmt"

// Node represents a node in the singly linked list
type Node struct {
    data int
    next *Node
}

// appendToStartOfTheList inserts a new node at the beginning of the list
func (head *Node) appendToStartOfTheList(data int) *Node {
    return &Node{data, head}
}

// addToTheEndOfTheList inserts a new node at the end of the list
func (head *Node) addToTheEndOfTheList(data int) *Node {
    if head == nil {
        return &Node{data, nil}
    }

    current := head
    for current.next != nil {
        current = current.next
    }
    current.next = &Node{data, nil}
    return head
}

// deleteAtBeginning removes the first node from the list
func (head *Node) deleteAtBeginning() *Node {
    if head == nil {
        fmt.Println("List is empty!")
        return nil
    }
    return head.next
}

// deleteAtEnd removes the last node from the list
func (head *Node) deleteAtEnd() *Node {
    if head == nil {
        fmt.Println("List is empty!")
        return nil
    }
    if head.next == nil {
        return nil
    }
    current := head
    for current.next.next != nil {
        current = current.next
    }
    current.next = nil
    return head
}

// lengthOfList returns the number of nodes in the list
func (head *Node) lengthOfList() int {
    count := 0
    current := head
    for current != nil {
        count++
        current = current.next
    }
    return count
}

// printList prints all the nodes in the list
func (head *Node) printList() {
    current := head
    for current != nil {
        fmt.Print(current.data, "->")
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    fmt.Println("Linked List Data Structure in GO Lang..!")

    var head *Node = nil
    for {
        fmt.Println("\nMenu:")
        fmt.Println("1. Insert node at beginning")
        fmt.Println("2. Insert node at end")
        fmt.Println("3. Deletion at the beginning")
        fmt.Println("4. Deletion at the end")
        fmt.Println("5. Length of the list")
        fmt.Println("0. Exit")
        fmt.Print("Enter your choice: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            fmt.Println("Enter the data to be inserted in the linked list")
            var inputData int
            fmt.Scanln(&inputData)
            head = head.appendToStartOfTheList(inputData)
            head.printList()
        case 2:
            fmt.Println("Enter the data to be inserted at the end of the list")
            var inputData int
            fmt.Scanln(&inputData)
            head = head.addToTheEndOfTheList(inputData)
            head.printList()
        case 3:
            head = head.deleteAtBeginning()
            head.printList()
        case 4:
            head = head.deleteAtEnd()
            head.printList()
        case 5:
            fmt.Println("Length of the list is:", head.lengthOfList())
            head.printList()
        case 0:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Try again.")
        }
    }
}