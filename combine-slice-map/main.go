package main

import "fmt"

func main() {
	data := make([]map[string]string, 0)

	data = append(data, map[string]string{"golang": "fulan", "java": "fulanah"}, map[string]string{"express": "marji"})

	for _, mapping := range data {
		for _, v := range mapping {
			fmt.Println(v)
		}
	}
}
