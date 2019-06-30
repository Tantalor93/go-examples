package main

import (
	"fmt"
	"sync"
)

func run(i int, group *sync.WaitGroup) {
	fmt.Printf("Run [%d]\n", i)
	group.Done()
}

func main() {
	fmt.Println("Starting")

	const numThreads = 10

	group := sync.WaitGroup{}
	group.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		go run(i, &group)
	}

	group.Wait()
	fmt.Println("Ending")
}
