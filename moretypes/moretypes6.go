package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Print(a[0], a[1])
	fmt.Print(a)

	primes := [6]int{2, 3, 7, 11, 13, 17}
	fmt.Print(primes)
}
