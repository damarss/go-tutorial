package main

import "fmt"

func main() {
	// array
	var a [4]int
	// a[0] = 2
	fmt.Println(a)

	b := [2]string{"Penn", "Teller"}
	fmt.Println(b)

	c := b
	fmt.Println(c)

	b[1] = "Kuda"
	fmt.Printf("b = %v | c = %v\n", b, c)


	// slice
	d := []string{"Penn", "Teller"}
	for i, v := range d {
		fmt.Printf("%v. %v\n", i+1, v)
	}
	d = append(d, "Kuda", "LMAO")
	fmt.Printf("d: %v\n", d)

}