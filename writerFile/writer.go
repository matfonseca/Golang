package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	textToWrite := []byte(os.Args[1])
	textToEdit := os.Args[1]

	//Create new file
	errorWrite := ioutil.WriteFile("example1.txt", textToWrite, 0777)
	showError(errorWrite)

	//Edit file
	file, errorOpenFile := os.OpenFile("./example.txt", os.O_APPEND, 0777)
	showError(errorOpenFile)
	writed, errorEdit := file.WriteString(textToEdit)
	showError(errorEdit)

	file.Close()

	fmt.Println(string(writed))

}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
