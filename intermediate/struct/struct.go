package main

import "fmt"

// Person struct represents a person with a name and age
type Person struct {
    Name string
    Age  int
}

// Greet returns a greeting message from the person (value receiver)
func (p Person) Greet() string {
    return p.Name + " says hello!"
}

// getNameAndAge increments the person's age and returns the new age (pointer receiver)
func (a *Person) getNameAndAge() int {
    a.Age++
    return a.Age
}

// Address struct represents a physical address
type Address struct {
    City    string
    Country string
}

// PhoneHomeCell struct represents home and cell phone numbers
type PhoneHomeCell struct {
    Home string
    Cell string
}

// Person2 struct demonstrates embedding other structs as named fields
type Person2 struct {
    Name          string
    Age           int
    address       Address
    PhoneHomeCell PhoneHomeCell
}

// Employee struct demonstrates an anonymous struct field (embedding)
// The Address field is embedded anonymously, so its fields are promoted
type Employee struct {
    Name string
    Age  int
    Address // Anonymous field (embedding)
}

func main() {
    // Create a Person instance
    p := Person{
        Name: "Alice",
        Age:  12,
    }

    fmt.Println(p) // Print the struct using default formatting
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

    // Anonymous struct example (struct type defined inline, used once)
    user := struct {
        Name string
        Age  int
    }{
        Name: "Bob",
        Age:  30,
    }
    fmt.Println(user)
    fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)

    // Call method with pointer receiver (Go handles address automatically)
    fmt.Println("After increment, Age:", p.getNameAndAge())
    fmt.Println("Greet:", p.Greet())
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

    // Create a Person2 instance with embedded Address and PhoneHomeCell structs (named fields)
    p2 := Person2{
        Name: "Rob",
        Age:  25,
        address: Address{
            City:    "New York",
            Country: "USA",
        },
        PhoneHomeCell: PhoneHomeCell{
            Home: "123-456-7890",
            Cell: "987-654-3210",
        },
    }

    fmt.Println(p2)
    fmt.Printf("Name: %s, Age: %d, City: %s, Country: %s\n",
        p2.Name, p2.Age, p2.address.City, p2.address.Country)

    // --- Anonymous Struct Field (Embedding) Example ---

    // Create an Employee instance with Address embedded anonymously
    emp := Employee{
        Name: "Charlie",
        Age:  40,
        Address: Address{
            City:    "London",
            Country: "UK",
        },
    }

    fmt.Println(emp)
    // Because Address is embedded anonymously, you can access its fields directly on Employee
    fmt.Printf("Employee Name: %s, Age: %d, City: %s, Country: %s\n",
        emp.Name, emp.Age, emp.City, emp.Country)

    // --- Struct Comparison Example ---
    // Comparing Address structs
    addr1 := Address{City: "London", Country: "UK"}
    addr2 := Address{City: "London", Country: "UK"}
    addr3 := Address{City: "Paris", Country: "France"}

    fmt.Println("addr1 == addr2:", addr1 == addr2) // true: all fields are equal
    fmt.Println("addr1 == addr3:", addr1 == addr3) // false: fields differ

    // Comparing Employee structs (with anonymous embedded Address)
    emp1 := Employee{
        Name: "Charlie",
        Age:  40,
        Address: Address{
            City:    "London",
            Country: "UK",
        },
    }
    emp2 := Employee{
        Name: "Charlie",
        Age:  40,
        Address: Address{
            City:    "London",
            Country: "UK",
        },
    }
    emp3 := Employee{
        Name: "Charlie",
        Age:  40,
        Address: Address{
            City:    "Paris",
            Country: "France",
        },
    }

    fmt.Println("emp1 == emp2:", emp1 == emp2) // true: all fields (including embedded) are equal
    fmt.Println("emp1 == emp3:", emp1 == emp3) // false: Address fields differ

    // Revision Notes:
    // - Anonymous struct: a struct defined and used inline, not as a named type.
    // - Anonymous struct field (embedding): a struct field declared with only the type, not a name.
    //   This "promotes" the embedded struct's fields and methods to the outer struct.
    // - You can access embedded fields directly (emp.City instead of emp.Address.City).
    // - Embedding is Go's way to achieve composition and reuse, similar to inheritance.
    // - Structs can be compared with == if all their fields are comparable.
}