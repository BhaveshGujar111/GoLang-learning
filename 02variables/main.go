package main

import "fmt"

// jwtToken := 300000 NOT ALLOWED
const LoginToken string = "jdnggdsnfk" // public

func main() {
	var username string = "bhavesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float64 = 255.242444443373848
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// default values and some aliases
	var anotherVariable float32
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
