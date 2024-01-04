package main

import "fmt"

func CalculateCost(carsCount int) uint {
	groups := int(carsCount / 10)
	individually := carsCount % 10

	total := uint((groups * 95000) + individually*10000)
	return total
}

func main() {
	fmt.Println(CalculateCost(37))
	fmt.Println(CalculateCost(21))
}
