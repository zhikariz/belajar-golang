package main

import "fmt"

func main() {
	var grupA = [...][3]string{[3]string{"a", "b", "c"}, [3]string{"d", "e", "f"}}
	var grupB = [...][3]string{{"g", "h", "i"}, {"j", "k", "l"}}

	fmt.Println(grupA)
	fmt.Println(grupB)
}
