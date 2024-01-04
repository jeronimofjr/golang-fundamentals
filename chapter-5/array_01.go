

import "fmt"

func array_01() {
	var x [5]int

	fmt.Println("Array que acabou de ser inicializado")
	fmt.Println(x)

	x[0] = 200
	x[1] = 300
	x[2] = 400
	x[3] = 500
	x[4] = 600

	fmt.Println("Array após atribuição de valores")
	fmt.Println(x)
}
