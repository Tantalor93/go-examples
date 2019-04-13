package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args) //argumenty
	fmt.Println("Argumenty bez jmena programu: ", os.Args[1:])
}
