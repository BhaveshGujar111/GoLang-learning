package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeder()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult, myMessage := proAdder(3, 4, 56, 7, 7, 5, 20)
	fmt.Println("PROResult is:", proResult)
	fmt.Println("PROMessage is:", myMessage)

}

func greeder() {
	fmt.Println("Namastey from golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi return another value"
}
