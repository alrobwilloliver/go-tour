package main

import (
	"fmt"
	"math"
)

func pow(x, y, lim float64) float64 {
	if x := math.Pow(x, y); x < lim {
		return x
	}
	return lim
}

func main () {
	fmt.Println(
		pow(2,3,10),
		pow(5,4, 36),
	)
}
