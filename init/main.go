package main

import "fmt"

var name string

func init() {
	fmt.Println("Hello from init")

	name = "Ondra"
}

func main() {
	fmt.Println("Hello from main to", name)

}
