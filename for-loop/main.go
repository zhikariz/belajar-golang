package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("data ke-", i)
		i++

		if i == 100 {
			break
		}
	}

	data := []string{"panji", "andrey", "syifa", "rahma"}

	for _, nama := range data {
		fmt.Println(nama)
	}

	for i := 1; i < 6; i++ {
		fmt.Println("ini i ke-", i)
	}
}
