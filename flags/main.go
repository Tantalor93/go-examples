package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "", "word to analyse")

	flag.Parse() // pred parsovanim musi predchazet deklarace flagu

	fmt.Println("provided word:", *wordPtr) // az pak se mohou pouzit
}
