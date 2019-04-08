package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "", "word to analyse")

	flag.Parse()

	fmt.Println("provided word:", *wordPtr)
}
