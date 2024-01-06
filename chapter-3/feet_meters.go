package main

import "fmt"

/*
Escreva um programa que converta pés em metros (1 pé = 0,3048 m).
*/
func main() {
	var feet float64
	const one_pe = 0.3048
	fmt.Print("Digite o valor em pés: ")
	fmt.Scanf("%f", &feet)

	meters := feet * one_pe

	fmt.Println(feet, "pés é igual a", meters, "metros")
}
