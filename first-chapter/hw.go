package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var Global int = 1234
var AnotherGlobal = -5678

func main() {
	fmt.Println("Hello World")
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Initial j value: ", j)
	j = Global
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d j=%d k=%.2f\n", Global, i, j, k)

	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}
	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("Zero")
	case "1":
		fmt.Println("One")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("Value: ", argument)
	}
	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int: ", argument)
	}
	switch {
	case value == 0:
		fmt.Println("Zero")
	case value < 0:
		fmt.Println("Negative")
	case value > 0:
		fmt.Println("Positive")
	default:
		fmt.Println("This should not happen: ", value)
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
	n := 0
	for ok := true; ok; ok = (n != 10) {
		fmt.Print(n*n, " ")
		n++
	}
	fmt.Println()

	n = 0
	for {
		if n == 10 {
			break
		}
		fmt.Print(n*n, " ")
		n++
	}
	fmt.Println()
	// This is a slice but range also works with arrays
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value: ", v)
	}

	counter := 100
	for i := 0; i <= counter; i++ {
		if i%2 == 0 {
			fmt.Println("Even number: ", i)
		}
	}

}
