package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal)

	smokeSignal <- false

	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
	attacked <- true
}

/*
	Explanation: ninja attack and taking more than 1 sec to complete
				main routine wait till the smokeSignal bool flag becomes true
				After that only main routine end their execution
				Each ninja require different channel because as we try to make
				// smokeSignal <- false \\ it becomes the deadlock condition.

	OUTPUT :
	A) without smokeSignal <- false
		Channels$ go run main.go
		Throwing ninja stars at Tommy
		true
		1.002420378s

	B) with smokeSignal <- false
		Channels$ go run main.go
		Throwing ninja stars at Tommy
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
        	/usr/local/src/golang-lco/Golang Dojo/Channels/main.go:18 +0xdb

		goroutine 18 [chan send]:
		main.attack({0x499a0b, 0x5}, 0x0?)
        	/usr/local/src/golang-lco/Golang Dojo/Channels/main.go:26 +0xa7
		created by main.main
        	/usr/local/src/golang-lco/Golang Dojo/Channels/main.go:16 +0xca
		exit status 2

*/
