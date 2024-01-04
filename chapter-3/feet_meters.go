package main

import "fmt"

func main() {
	var feet float64
	fmt.Print("Digite o valor em pés: ")
	fmt.Scanf("%f", &feet)

	meters := feet * 0.3048

	fmt.Println(feet, "pés é igual a", meters, "metros")
}
