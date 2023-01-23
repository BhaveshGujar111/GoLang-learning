package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This need to go in a file - LCO.in"
	file, err := os.Create("./mylcofofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	len, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is:", len)
	defer file.Close()
	readFile("./mylcofofile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
