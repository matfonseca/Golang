package main

import (
	"fmt"
	"strconv"
)

func main() {
	var result int = sum(1, 3)
	fmt.Println("El resultado es :" + strconv.Itoa(result))
}

func sum(a int, b int) int {
	return a + b
}
