package main

func sumOfNumbers(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumOfNumbers(n-1)
}

