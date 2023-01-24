package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	// Go-Routines - Concurrently running routines
	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}
	for _, evilNinga := range evilNinjas {
		go attack(evilNinga)
	}

	// main routines executes concurrently with attack go-routines,
	// instead program stops it execution when main routine ends, that is why there is requirement to channels

	time.Sleep(time.Second)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second)
}

/*
This is the output because of the sleep function
Else program prints Unexpected output when main routine ends

	Concurrency$ go run main.go
	Throwing ninja stars at Andy
	Throwing ninja stars at Johnny
	Throwing ninja stars at Bobby
	Throwing ninja stars at Tommy
	1.01391582s
*/
