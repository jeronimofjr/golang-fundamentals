package main

import "fmt"

/* escreva um programa que converta
de Fahrenheit para Celsius (C = (F − 32) * 5/9)
*/
func main() {
	var fahrenheit float64
	fmt.Print("Digite o valor em graus Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)

	C := ((fahrenheit - 32) * 5 / 9)

	fmt.Println(fahrenheit, "graus em fahrenheit é igual a", C, "graus em celsius")
}
