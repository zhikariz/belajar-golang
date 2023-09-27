package main

import "fmt"

func main() {
	var fruits = make([]string, 4)
	fruits[0] = "apple"
	fruits[1] = "grape"
	fruits[2] = "banana"
	fruits[3] = "melon"

	// for i := 0; i < len(fruits); i++ {
	// fmt.Printf("element ke %d adalah %s\n", i, fruits[i])
	// }

	for _, value := range fruits {
		fmt.Printf("elementnya adalah %s\n", value)
	}
}
