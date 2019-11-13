package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	peliculas := []string{"volver al futuro", "rapido y furioso", "titanic"}

	for _, peli := range peliculas {
		fmt.Println(peli)
	}

	num := 30
	switch num {
	case 10:
		println("es 10")
	default:
		println("no es 10")
	}
}
