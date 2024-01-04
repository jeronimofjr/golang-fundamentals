package main

import "fmt"

func main() {
	x := []float64{4.35, 2.22, 00.1, 3.25}
	vazio := []float64{}
	pis := [3]float64{3.14, 3.14, 3.14}

	fmt.Println(x)
	x[0] = 0.5
	fmt.Println(x)
	fmt.Println(vazio)
	fmt.Println(pis)
	// vazio[0] = 3
	// fmt.Println(vazio)
}
