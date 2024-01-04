package main

import "fmt"

func main() {
	var i int
	fmt.Print("Digite um número, por favor: ")

	fmt.Scan(&i)

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 3:
		fmt.Println("Três")
	case 4:
		fmt.Println("Quatro")
	case 5:
		fmt.Println("Cinco")
	case 6:
		fmt.Println("Seis")
	case 7:
		fmt.Println("Sete")
	default:
		fmt.Println("Valor desconhecido")
	}
}
