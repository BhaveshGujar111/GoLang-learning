package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// Mostly Used
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at LCO.in")

}
