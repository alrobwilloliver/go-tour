package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["St James Park"] = Vertex{
		50.730713888889, -3.52115,
	}
	fmt.Print(m)
}
