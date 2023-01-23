package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// fruitList[2] = "Apple"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list id: ", fruitList)
	fmt.Println("Fruit length id: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)

}
