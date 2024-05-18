package main

import "fmt"

func main() {
	// input x
	var x int
	fmt.Print("Enter x : ")
	fmt.Scan(&x)

	// input y
	var y int
	fmt.Print("Enter y: ")
	fmt.Scan(&y)

	// input position
	var position int
	fmt.Print("Enter position: ")
	fmt.Scan(&position)

	var modResult = x % y
	var decimalPosition []int

	// loop position times
	for i := 0; i < position; i++ {
		modResult = modResult * 10
		var digit = modResult / y
		// apend to slice
		decimalPosition = append(decimalPosition, digit)
		modResult = modResult % y
	}
	fmt.Print("Decimal position is: ")
	fmt.Print(decimalPosition[position-1])
}
