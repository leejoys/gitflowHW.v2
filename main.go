package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Println("Hello World!")
}

func main() {
	greet()
	fmt.Printf("Time is: %s", time.Now())
}
