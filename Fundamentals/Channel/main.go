package main

import (
	"fmt"
	"math/rand/v2"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)
	intChan <- randomNumber
}

func RandomNumber(n int) int {
	value := rand.IntN(n)

	return value
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	fmt.Println("Hello world")

	num := <-intChan
	fmt.Println(num)
}
