package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	arr := []int{0,1}
	return func() int {
		arr = append(arr[:0], arr[1], arr[0]+arr[1])
		return arr[len(arr) - 1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
