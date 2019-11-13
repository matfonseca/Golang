package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	println("Los argumentos de la consola son:")
	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		edad, _ := strconv.Atoi(os.Args[1])

		if edad > 10 {
			println("la edad es mayor a 10")
		} else if edad == 10 {
			println("la edad es 10")
		} else {
			println("la edad es menor a 10")
		}
	}

}
