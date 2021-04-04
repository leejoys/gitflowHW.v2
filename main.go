package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Printf("Hello Word!")
}

func main() {
	greet()
	fmt.Printf("Time is: %s", time.Now())
}
