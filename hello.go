package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	// call function
	fmt.Println(add(42, 13))
	fmt.Println(add(rand.Intn(100), rand.Intn(100)))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
