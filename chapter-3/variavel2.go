package main

import "fmt"

func main() {
	fmt.Print("Entre com um número: ")
	var input float64
	const pi float64 = 3.14
	var data = 4.5
	var x string = "opa"
	et := 2.62
	fmt.Scanf("%f", &input)
	fmt.Println("Pi vezes o valor que vc digitou é igual a: ", input*pi)
	fmt.Println(et)
	fmt.Println(data)
}
