package main

import (
	"fmt"
)

func main() {
	imprimirMessage("hola", "mundo", "!")
}

func imprimirMessage(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}
