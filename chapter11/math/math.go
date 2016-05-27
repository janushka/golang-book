package math

import "fmt"

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func TellSomething() string {
	return tellAnotherThing()
}

func tellAnotherThing() string {
	x := 5
	zero(&x)
	fmt.Println(x)
	return "Bla2 Bla2 Bla2..."
}

func zero(xPtr *int) {
	*xPtr = 0
}