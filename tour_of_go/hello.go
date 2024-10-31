package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

var c, python, java bool

// Variables with initializers
var j, k int = 1, 2

// Basic types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// START Numeric Constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// END Numeric Constants

// IF
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprintln(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

// Exercise: Loops and Functions

// Switch

func main() {
	fmt.Printf("Hello, Thien")

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("Now you have %g problems .\n", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(20))

	var i int
	fmt.Println(i, c, python, java)

	// Variables with initializers
	var c, python, java = true, false, "no!"
	fmt.Println(j, k, c, python, java)

	// Basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) // T = type, v = value
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Type inference
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// For
	sum := 0
	for o := 0; o < 10; o++ {
		sum += o
	}
	fmt.Println(sum)

	sumTemp := 0
	for q := 0; q < 5; q++ {
		sumTemp += q
	}
	fmt.Println("sum temp:", sumTemp)

	// For continued
	sum = 50
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// IF
	fmt.Println(sqrt(2), sqrt(-4))

	// If with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// SWITCH
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
