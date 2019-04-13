package main

import "fmt"

func main() {
	fmt.Println(`ahoj\n
	petre`) //raw string literal

	fmt.Println("ahoj\n" + //interpreted
		"Petre")
}
