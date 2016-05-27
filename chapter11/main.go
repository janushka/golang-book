package main

import "fmt"

/*func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)

	fmt.Println(avg)
	fmt.Println(math.TellSomething())
}*/

var name = "Salomé"

func main() {
	first := "Salomé"
	defer func() {
		str := recover()
		fmt.Scanln(str + name + first)
	}()
	panic("ALARM!!")
}