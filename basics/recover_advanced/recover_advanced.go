// Package main demonstrates advanced usage of Go's recover mechanism
package main

import (
	"fmt"
	"runtime/debug"
)

// CustomError represents an application-specific error
type CustomError struct {
	Message string
}

// Error implements the error interface
func (e CustomError) Error() string {
	return e.Message
}

// recoverWithStackTrace demonstrates recover with stack trace printing
func recoverWithStackTrace() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
			fmt.Println("\nStack trace:")
			fmt.Println(string(debug.Stack()))
		}
	}()

	// Cause a panic
	panic("something went wrong")
}

// recoverAndRethrow demonstrates how to recover and then re-panic
// if needed (useful for cleanup but still propagating the panic)
func recoverAndRethrow() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Cleanup after panic...")
			
			// Re-panic with the same value
			panic(r)
		}
	}()

	panic("critical error")
}

// panicWithCustomError demonstrates panicking with a custom error type
func panicWithCustomError() {
	panic(CustomError{Message: "custom error occurred"})
}

// recoverCustomError demonstrates type checking in recover
func recoverCustomError() {
	defer func() {
		if r := recover(); r != nil {
			// Type assertion to check if it's our custom error
			if customErr, ok := r.(CustomError); ok {
				fmt.Println("Recovered from custom error:", customErr.Message)
			} else if err, ok := r.(error); ok {
				fmt.Println("Recovered from standard error:", err.Error())
			} else {
				fmt.Println("Recovered from unknown panic:", r)
			}
		}
	}()

	panicWithCustomError()
}

// nestedPanic demonstrates panic inside a recover function
func nestedPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Outer recover caught:", r)
			
			// Set up another recover for the nested panic
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Inner recover caught:", r)
				}
			}()
			
			// This will cause another panic
			panic("nested panic")
		}
	}()

	panic("original panic")
}

// recoverFromGoroutine shows that panics don't cross goroutine boundaries
func recoverFromGoroutine() {
	// Set up a recover in the main goroutine
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("This won't execute because the panic is in another goroutine")
		}
	}()

	// Create a goroutine that will panic
	go func() {
		// Each goroutine needs its own recover
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic in goroutine:", r)
			}
		}()
		
		panic("panic in goroutine")
	}()

	// Give the goroutine time to execute
	fmt.Println("Waiting for goroutine...")
	// In a real program, you would use proper synchronization
	// This is just for demonstration
	for i := 0; i < 1000000; i++ {
		// Busy wait
	}
}

// recoverWithErrorHandling shows how to convert panics to errors
func recoverWithErrorHandling() (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Convert the panic to an error return
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	// This will panic
	panic("something failed")
}

func main() {
	fmt.Println("=== Advanced Recover Examples ===")

	fmt.Println("\n1. Recover with stack trace:")
	recoverWithStackTrace()

	fmt.Println("\n2. Recover custom error types:")
	recoverCustomError()

	fmt.Println("\n3. Recover in goroutine:")
	recoverFromGoroutine()

	fmt.Println("\n4. Convert panic to error:")
	if err := recoverWithErrorHandling(); err != nil {
		fmt.Println("Function returned error:", err)
	}

	fmt.Println("\n5. Nested panic and recover:")
	nestedPanic()

	fmt.Println("\n6. Recover and rethrow (this will terminate the program):")
	// This is commented out because it would terminate the program
	// Uncomment to see the effect
	// recoverAndRethrow()

	fmt.Println("\nProgram completed")
}