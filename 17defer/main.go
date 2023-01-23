package main

import "fmt"

func main() {
	defer fmt.Println("Hello World 1")
	defer fmt.Println("Hello World 2")
	defer fmt.Println("Hello World 3")
	defer fmt.Println("Hello World 4")
	fmt.Println("Defer in Golang")
	myDefer()

}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Print(i)
	}
}
