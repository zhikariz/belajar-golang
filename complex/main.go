package main

import (
	"fmt"
)

func main() {
	c1 := complex(2, 3)
	fmt.Println(c1)
	c2 := 4 + 5i // complex initializer syntax a + ib
	fmt.Println(c2)
	c3 := c1 + c2            // addition just like other variables
	fmt.Println("Add: ", c3) // prints "Add: (6+8i)"
	re := real(c3)           // get real part
	im := imag(c3)           // get imaginary part
	fmt.Println(re, im)      // prints 6 8
}
