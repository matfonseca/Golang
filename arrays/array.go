package main

import (
	"fmt"
)

func main() {
	var array [3]string
	array[0] = "hola"
	array[1] = "mundo"
	array[2] = "!"

	fmt.Println(array)

	array2 := []string{"hola", "mundo", "!"}
	fmt.Println(array2)

	array2 = append(array2, "!")
	fmt.Println(array2)

	fmt.Println(array2[0:1])
}
