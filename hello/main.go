package main

import (
	"fmt"
	"github.com/tantalor93/go-examples/hello/geometry"
	"github.com/tantalor93/go-examples/stringutil"
)

func main() {
	square := geometry.Square{A: geometry.Vertex{X: 0, Y:0}, B: geometry.Vertex{X: 1, Y: 1}}
	fmt.Println(square.Perimeter())

	const hello = "hello, world\n"
	const a = 1;
	if a < 2 {
		fmt.Println("ahoj")
	}
	defer fmt.Println("ahojky")
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 25
	fmt.Println(i)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	for j, v := range primes {
		v = 1
		fmt.Printf("2**%d = %d\n", j, v)
	}
	fmt.Println("primes:", primes)
	primes[1:][0] = 5
	fmt.Println("primes:", primes[1:])
	fmt.Printf(stringutil.Reverse(hello))
}
