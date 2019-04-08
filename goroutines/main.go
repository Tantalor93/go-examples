package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go randomSleepAndPrint()
	go randomSleepAndPrint()
	time.Sleep(time.Second * 10)
	fmt.Println("konec")

	channel := make(chan int)

	ints := []int{1, 3, 5, 7, 9, 11, 3, 5, 4, 2}

	go sum(ints[:len(ints)/2], channel)
	go sum(ints[len(ints)/2:], channel)

	a, b := <-channel, <-channel
	fmt.Println(a + b)
}

func sum(arr []int, ch chan int) {
	var s = 0
	for _, v := range arr {
		s += v
	}
	ch <- s
}

func randomSleepAndPrint() {
	i := time.Duration(rand.Intn(10))
	duration := time.Second * i
	time.Sleep(duration)
	fmt.Println("AHOJ")
}
