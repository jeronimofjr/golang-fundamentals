package main

import "fmt"

func main() {
	var n1 int
	var n2 int

	fmt.Scanln(&n1, &n2)

	if n1 == n2 {
		output := fmt.Sprintf("%d e %d são iguais", n1, n2)

		fmt.Println(output)
	} else {
		output := fmt.Sprintf("%d e %d não são iguais", n1, n2)

		fmt.Println(output)
	}
}
