package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int) //dynamicka alokace mapy
	m["A"] = 1
	fmt.Println(m)

	var a [5]int //staticka alokace , defaultne vsechno nula
	a[0] = 1

	ints := make([]int, 10) //dynamicka alokace slice
	ints[0] = 1

	fmt.Println(len(ints))
}
