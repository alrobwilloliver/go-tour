package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0;
	var count = 0
	for ; count < 10; count++ {
		var prev = z
		z -= (z*z - x) / (2*z)
		if prev - 1 == z {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}
