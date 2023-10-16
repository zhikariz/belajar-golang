package main

import "fmt"

func main() {
	chicken := make(map[string]int)

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	var data = map[string]int{
		"two":   2,
		"three": 3,
	}
	data["one"] = 1
	fmt.Println(data)

	// delete(data, "three")

	for key, value := range data {
		fmt.Printf("Keynya adalah %s Valuenya adalah %d\n", key, value)
	}

	if value, isExist := data["three"]; !isExist {
		fmt.Println("that data was not exists")
	} else {
		fmt.Println(value)
	}
}
