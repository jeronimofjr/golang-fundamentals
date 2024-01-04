package main

import "fmt"

func main() {
	nome := "jeronimo"
	nomes := []string{"Jeronimo", "Salah", "Junior"}

	fmt.Println("PERCORRENDO UM NOME")
	for idx, value := range nome {
		fmt.Println("Posição: ", idx)
		fmt.Println("Letra:", string(value))
	}

	fmt.Println("PERCORRENDO UMA LISTA DE NOMES")
	for _, value := range nomes {
		fmt.Println(value)
	}
}
