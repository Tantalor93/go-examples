package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Argumenty bez jmena programu: ", os.Args[1:])
	fmt.Printf("hello, world2\n")
	fmt.Println(fmt.Sprint(1, "."))

	m := make(map[string]int)
	m["A"] = 1
	fmt.Println(m)

	ints := make([]int, 10)
	ints[0] = 1
}
