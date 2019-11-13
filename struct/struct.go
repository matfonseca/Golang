package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Nombre       string
	Apellido     string
	Edad         int
	Nacionalidad string
}

func main() {
	var person1 = Person{
		Nombre:       "Matias",
		Apellido:     "Fonseca",
		Edad:         23,
		Nacionalidad: "ARG"}

	person1.Nombre = "Jorge"

	fmt.Println("Se ingreso la persona con el nombre: " + person1.Nombre + " y el apellido: " + person1.Apellido + ", que tiene la edad de: " + strconv.Itoa(person1.Edad) + " Y su nacionalidad es: " + person1.Nacionalidad)
}
