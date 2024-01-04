package main

import "fmt"

func main() {
	x := make(map[string]string)

	x["key"] = "sucesso"

	if name, ok := x["key"]; ok {
		fmt.Println(name, ok)
	}
}
