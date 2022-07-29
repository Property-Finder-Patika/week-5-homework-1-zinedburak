package main

import "fmt"

var MustBeZero int

// run the following code with go run -race main.go to see the error
func main() {
	MustBeZero = 10
	// Most simple race condition Both of them are trying to update the same value
	go add()
	go add()
	fmt.Println(MustBeZero)
}

func add() {
	MustBeZero++
}
