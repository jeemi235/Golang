package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcomr to files")
	content := "this needs to go in the file"

	file, err := os.Create("./mylco.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is:", length)
	defer file.Close()
	readfile("./mylco.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
