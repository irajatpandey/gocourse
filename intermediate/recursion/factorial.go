package main

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	smallOutput := factorial(n-1)
	return smallOutput * n
	
}