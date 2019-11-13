package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	text, errorFile := ioutil.ReadFile("example.txt")

	showError(errorFile)

	fmt.Println(string(text))

}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
