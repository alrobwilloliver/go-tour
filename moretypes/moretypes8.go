package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]
	fmt.Print(a, b)

	b[0] = "XXX"
	fmt.Print(a, b, names)
}
