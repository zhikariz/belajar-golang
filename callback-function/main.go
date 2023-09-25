package main

import "fmt"

func callbackExample(x int, callback func(int)) {
	callback(x)
}

func main() {
	callbackExample(10, func(x int) {
		fmt.Println("Callback called with:", x)
	})
}
