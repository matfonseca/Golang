package main

import (
	"fmt"
	"strconv"
)

func main() {
	var result int = getPrecio(3, 2)
	fmt.Println("El precio es: " + strconv.Itoa(result))
}

func getPrecio(n1 int, n2 int) int {
	precio := func() int {
		return n1 * n2
	}
	return precio()
}
