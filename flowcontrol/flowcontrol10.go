package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Printf("It's today!")
	case today + 1:
		fmt.Printf("It's tomorrow!")
	case today + 2:
		fmt.Printf("It's in 2 days!")
	default:
		fmt.Printf("It's too far away!")
	}
}
