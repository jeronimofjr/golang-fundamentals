package main

import "fmt"

func main() {
	var x [5]float64

	x[0] = 13
	x[1] = 10
	x[2] = 7
	x[3] = 3
	x[4] = 43

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
