package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Print("The Value: ", m["Answer"], "\n")

	m["Answer"] = 50
	fmt.Print("The Value: ", m["Answer"], "\n")

	delete(m, "Answer")
	fmt.Println("The Value ", m["Answer"], "\n")

	v, ok := m["Answer"]
	fmt.Println("The Value ", v, "Present?", ok)
}
