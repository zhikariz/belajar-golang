package main

import "fmt"

func factorial(n int, i int) int {
	fmt.Printf("panggilan ke %d dengan n nilainya adalah %d\n", i, n)
	i++
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1, i)
}

func main() {
	result := factorial(6, 1)
	fmt.Println(result) // Output: 120
}
